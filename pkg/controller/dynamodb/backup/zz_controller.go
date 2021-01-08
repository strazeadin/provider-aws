/*
Copyright 2020 The Crossplane Authors.

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

package backup

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/dynamodb"
	svcsdk "github.com/aws/aws-sdk-go/service/dynamodb"
	svcsdkapi "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/dynamodb/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an Backup resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Backup in AWS"
	errDescribe      = "failed to describe Backup"
	errDelete        = "failed to delete Backup"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Backup)
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
	cr, ok := mg.(*svcapitypes.Backup)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeBackupInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeBackupWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errors.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateBackup(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        e.isUpToDate(cr, resp),
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.Backup)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateBackupInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateBackupWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreate)
	}

	if resp.BackupDetails.BackupArn != nil {
		cr.Status.AtProvider.BackupARN = resp.BackupDetails.BackupArn
	}
	if resp.BackupDetails.BackupCreationDateTime != nil {
		cr.Status.AtProvider.BackupCreationDateTime = &metav1.Time{*resp.BackupDetails.BackupCreationDateTime}
	}
	if resp.BackupDetails.BackupExpiryDateTime != nil {
		cr.Status.AtProvider.BackupExpiryDateTime = &metav1.Time{*resp.BackupDetails.BackupExpiryDateTime}
	}
	if resp.BackupDetails.BackupSizeBytes != nil {
		cr.Status.AtProvider.BackupSizeBytes = resp.BackupDetails.BackupSizeBytes
	}
	if resp.BackupDetails.BackupStatus != nil {
		cr.Status.AtProvider.BackupStatus = resp.BackupDetails.BackupStatus
	}
	if resp.BackupDetails.BackupType != nil {
		cr.Status.AtProvider.BackupType = resp.BackupDetails.BackupType
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	return e.update(ctx, mg)

}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Backup)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteBackupInput(cr)
	if err := e.preDelete(ctx, cr, input); err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	_, err := e.client.DeleteBackupWithContext(ctx, input)
	return errors.Wrap(cpresource.Ignore(IsNotFound, err), errDelete)
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.DynamoDBAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		update:         nopUpdate,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.DynamoDBAPI
	preObserve     func(context.Context, *svcapitypes.Backup, *svcsdk.DescribeBackupInput) error
	postObserve    func(context.Context, *svcapitypes.Backup, *svcsdk.DescribeBackupOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.BackupParameters, *svcsdk.DescribeBackupOutput) error
	isUpToDate     func(*svcapitypes.Backup, *svcsdk.DescribeBackupOutput) bool
	preCreate      func(context.Context, *svcapitypes.Backup, *svcsdk.CreateBackupInput) error
	postCreate     func(context.Context, *svcapitypes.Backup, *svcsdk.CreateBackupOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Backup, *svcsdk.DeleteBackupInput) error
	// TODO(muvaf): update is not supported in most of the resources. this is temporary.
	update func(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Backup, *svcsdk.DescribeBackupInput) error {
	return nil
}
func nopPostObserve(context.Context, *svcapitypes.Backup, *svcsdk.DescribeBackupOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error) {
	return managed.ExternalObservation{}, nil
}
func nopLateInitialize(*svcapitypes.BackupParameters, *svcsdk.DescribeBackupOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.Backup, *svcsdk.DescribeBackupOutput) bool {
	return true
}

func nopPreCreate(context.Context, *svcapitypes.Backup, *svcsdk.CreateBackupInput) error {
	return nil
}
func nopPostCreate(context.Context, *svcapitypes.Backup, *svcsdk.CreateBackupOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error) {
	return managed.ExternalCreation{}, nil
}
func nopPreDelete(context.Context, *svcapitypes.Backup, *svcsdk.DeleteBackupInput) error {
	return nil
}
func nopUpdate(context.Context, cpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
