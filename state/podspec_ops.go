// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"github.com/juju/errors"
	"gopkg.in/juju/charm.v6"
	"gopkg.in/juju/names.v3"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"

	"github.com/juju/juju/core/leadership"
)

type setPodSpecOperation struct {
	m      *CAASModel
	appTag names.ApplicationTag
	spec   *string

	tokenAwareTxnBuilder func(int) ([]txn.Op, error)
}

// newSetPodSpecOperation returns a ModelOperation for updating the PodSpec
// for a particular application. A nil token can be specified to bypass the
// leadership check.
func newSetPodSpecOperation(model *CAASModel, token leadership.Token, appTag names.ApplicationTag, spec *string) *setPodSpecOperation {
	op := &setPodSpecOperation{
		m:      model,
		appTag: appTag,
		spec:   spec,
	}

	if token != nil {
		op.tokenAwareTxnBuilder = buildTxnWithLeadership(op.buildTxn, token)
	}
	return op
}

// Build implements ModelOperation.
func (op *setPodSpecOperation) Build(attempt int) ([]txn.Op, error) {
	if op.tokenAwareTxnBuilder != nil {
		return op.tokenAwareTxnBuilder(attempt)
	}
	return op.buildTxn(attempt)
}

func (op *setPodSpecOperation) buildTxn(_ int) ([]txn.Op, error) {
	var prereqOps []txn.Op
	appTagID := op.appTag.Id()
	app, err := op.m.State().Application(appTagID)
	if err != nil {
		return nil, errors.Annotate(err, "setting pod spec")
	}
	if app.Life() != Alive {
		return nil, errors.Annotate(
			errors.Errorf("application %s not alive", app.String()),
			"setting pod spec",
		)
	}
	// The app's charm may not be there yet (as is the case when migrating).
	// This check is for checking the k8s-spec-set call.
	ch, _, err := app.Charm()
	if err != nil && !errors.IsNotFound(err) {
		return nil, errors.Trace(err)
	} else if err == nil {
		if ch.Meta().Deployment != nil && ch.Meta().Deployment.DeploymentMode == charm.ModeOperator {
			return nil, errors.New("cannot set k8s spec on an operator charm")
		}
	}
	prereqOps = append(prereqOps, txn.Op{
		C:      applicationsC,
		Id:     app.doc.DocID,
		Assert: isAliveDoc,
	})

	sop := txn.Op{
		C:  podSpecsC,
		Id: applicationGlobalKey(appTagID),
	}
	existing, err := op.m.podInfo(op.appTag)
	if err == nil {
		updates := bson.D{{"$inc", bson.D{{"upgrade-counter", 1}}}}
		if op.spec != nil {
			updates = append(updates, bson.DocElem{"$set", bson.D{{"spec", *op.spec}}})
		}
		sop.Assert = bson.D{{"upgrade-counter", existing.UpgradeCounter}}
		sop.Update = updates
	} else if errors.IsNotFound(err) {
		sop.Assert = txn.DocMissing
		var specStr string
		if op.spec != nil {
			specStr = *op.spec
		}
		sop.Insert = containerSpecDoc{Spec: specStr}
	} else {
		return nil, errors.Annotate(err, "setting pod spec")
	}
	return append(prereqOps, sop), nil
}

// Done implements ModelOperation.
func (op *setPodSpecOperation) Done(err error) error { return err }
