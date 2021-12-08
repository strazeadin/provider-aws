/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package integrationresponse

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an IntegrationResponse resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create IntegrationResponse in AWS"
	errUpdate        = "cannot update IntegrationResponse in AWS"
	errDescribe      = "failed to describe IntegrationResponse"
	errDelete        = "failed to delete IntegrationResponse"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.IntegrationResponse)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.IntegrationResponse)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetIntegrationResponseInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetIntegrationResponseWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateIntegrationResponse(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.IntegrationResponse)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateIntegrationResponseInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateIntegrationResponseWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.ContentHandlingStrategy != nil {
		cr.Spec.ForProvider.ContentHandlingStrategy = resp.ContentHandlingStrategy
	} else {
		cr.Spec.ForProvider.ContentHandlingStrategy = nil
	}
	if resp.IntegrationResponseId != nil {
		cr.Status.AtProvider.IntegrationResponseID = resp.IntegrationResponseId
	} else {
		cr.Status.AtProvider.IntegrationResponseID = nil
	}
	if resp.IntegrationResponseKey != nil {
		cr.Spec.ForProvider.IntegrationResponseKey = resp.IntegrationResponseKey
	} else {
		cr.Spec.ForProvider.IntegrationResponseKey = nil
	}
	if resp.ResponseParameters != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range resp.ResponseParameters {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		cr.Spec.ForProvider.ResponseParameters = f3
	} else {
		cr.Spec.ForProvider.ResponseParameters = nil
	}
	if resp.ResponseTemplates != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range resp.ResponseTemplates {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		cr.Spec.ForProvider.ResponseTemplates = f4
	} else {
		cr.Spec.ForProvider.ResponseTemplates = nil
	}
	if resp.TemplateSelectionExpression != nil {
		cr.Spec.ForProvider.TemplateSelectionExpression = resp.TemplateSelectionExpression
	} else {
		cr.Spec.ForProvider.TemplateSelectionExpression = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.IntegrationResponse)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateIntegrationResponseInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateIntegrationResponseWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.IntegrationResponse)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteIntegrationResponseInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteIntegrationResponseWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ApiGatewayV2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.ApiGatewayV2API
	preObserve     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.GetIntegrationResponseInput) error
	postObserve    func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.GetIntegrationResponseOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.IntegrationResponseParameters, *svcsdk.GetIntegrationResponseOutput) error
	isUpToDate     func(*svcapitypes.IntegrationResponse, *svcsdk.GetIntegrationResponseOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.CreateIntegrationResponseInput) error
	postCreate     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.CreateIntegrationResponseOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.DeleteIntegrationResponseInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.DeleteIntegrationResponseOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.UpdateIntegrationResponseInput) error
	postUpdate     func(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.UpdateIntegrationResponseOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.GetIntegrationResponseInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.IntegrationResponse, _ *svcsdk.GetIntegrationResponseOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.IntegrationResponseParameters, *svcsdk.GetIntegrationResponseOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.IntegrationResponse, *svcsdk.GetIntegrationResponseOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.CreateIntegrationResponseInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.IntegrationResponse, _ *svcsdk.CreateIntegrationResponseOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.DeleteIntegrationResponseInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.IntegrationResponse, _ *svcsdk.DeleteIntegrationResponseOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.IntegrationResponse, *svcsdk.UpdateIntegrationResponseInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.IntegrationResponse, _ *svcsdk.UpdateIntegrationResponseOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
