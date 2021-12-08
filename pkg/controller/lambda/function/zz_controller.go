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

package function

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/lambda"
	svcsdk "github.com/aws/aws-sdk-go/service/lambda"
	svcsdkapi "github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane/provider-aws/apis/lambda/v1alpha1"
	awsclient "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an Function resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Function in AWS"
	errUpdate        = "cannot update Function in AWS"
	errDescribe      = "failed to describe Function"
	errDelete        = "failed to delete Function"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Function)
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
	cr, ok := mg.(*svcapitypes.Function)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetFunctionInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetFunctionWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateFunction(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.Function)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateFunctionInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateFunctionWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.CodeSha256 != nil {
		cr.Status.AtProvider.CodeSHA256 = resp.CodeSha256
	} else {
		cr.Status.AtProvider.CodeSHA256 = nil
	}
	if resp.CodeSize != nil {
		cr.Status.AtProvider.CodeSize = resp.CodeSize
	} else {
		cr.Status.AtProvider.CodeSize = nil
	}
	if resp.DeadLetterConfig != nil {
		f2 := &svcapitypes.DeadLetterConfig{}
		if resp.DeadLetterConfig.TargetArn != nil {
			f2.TargetARN = resp.DeadLetterConfig.TargetArn
		}
		cr.Spec.ForProvider.DeadLetterConfig = f2
	} else {
		cr.Spec.ForProvider.DeadLetterConfig = nil
	}
	if resp.Description != nil {
		cr.Spec.ForProvider.Description = resp.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Environment != nil {
		f4 := &svcapitypes.Environment{}
		if resp.Environment.Variables != nil {
			f4f1 := map[string]*string{}
			for f4f1key, f4f1valiter := range resp.Environment.Variables {
				var f4f1val string
				f4f1val = *f4f1valiter
				f4f1[f4f1key] = &f4f1val
			}
			f4.Variables = f4f1
		}
		cr.Spec.ForProvider.Environment = f4
	} else {
		cr.Spec.ForProvider.Environment = nil
	}
	if resp.FileSystemConfigs != nil {
		f5 := []*svcapitypes.FileSystemConfig{}
		for _, f5iter := range resp.FileSystemConfigs {
			f5elem := &svcapitypes.FileSystemConfig{}
			if f5iter.Arn != nil {
				f5elem.ARN = f5iter.Arn
			}
			if f5iter.LocalMountPath != nil {
				f5elem.LocalMountPath = f5iter.LocalMountPath
			}
			f5 = append(f5, f5elem)
		}
		cr.Spec.ForProvider.FileSystemConfigs = f5
	} else {
		cr.Spec.ForProvider.FileSystemConfigs = nil
	}
	if resp.FunctionArn != nil {
		cr.Status.AtProvider.FunctionARN = resp.FunctionArn
	} else {
		cr.Status.AtProvider.FunctionARN = nil
	}
	if resp.FunctionName != nil {
		cr.Status.AtProvider.FunctionName = resp.FunctionName
	} else {
		cr.Status.AtProvider.FunctionName = nil
	}
	if resp.Handler != nil {
		cr.Spec.ForProvider.Handler = resp.Handler
	} else {
		cr.Spec.ForProvider.Handler = nil
	}
	if resp.ImageConfigResponse != nil {
		f9 := &svcapitypes.ImageConfigResponse{}
		if resp.ImageConfigResponse.Error != nil {
			f9f0 := &svcapitypes.ImageConfigError{}
			if resp.ImageConfigResponse.Error.ErrorCode != nil {
				f9f0.ErrorCode = resp.ImageConfigResponse.Error.ErrorCode
			}
			if resp.ImageConfigResponse.Error.Message != nil {
				f9f0.Message = resp.ImageConfigResponse.Error.Message
			}
			f9.Error = f9f0
		}
		if resp.ImageConfigResponse.ImageConfig != nil {
			f9f1 := &svcapitypes.ImageConfig{}
			if resp.ImageConfigResponse.ImageConfig.Command != nil {
				f9f1f0 := []*string{}
				for _, f9f1f0iter := range resp.ImageConfigResponse.ImageConfig.Command {
					var f9f1f0elem string
					f9f1f0elem = *f9f1f0iter
					f9f1f0 = append(f9f1f0, &f9f1f0elem)
				}
				f9f1.Command = f9f1f0
			}
			if resp.ImageConfigResponse.ImageConfig.EntryPoint != nil {
				f9f1f1 := []*string{}
				for _, f9f1f1iter := range resp.ImageConfigResponse.ImageConfig.EntryPoint {
					var f9f1f1elem string
					f9f1f1elem = *f9f1f1iter
					f9f1f1 = append(f9f1f1, &f9f1f1elem)
				}
				f9f1.EntryPoint = f9f1f1
			}
			if resp.ImageConfigResponse.ImageConfig.WorkingDirectory != nil {
				f9f1.WorkingDirectory = resp.ImageConfigResponse.ImageConfig.WorkingDirectory
			}
			f9.ImageConfig = f9f1
		}
		cr.Status.AtProvider.ImageConfigResponse = f9
	} else {
		cr.Status.AtProvider.ImageConfigResponse = nil
	}
	if resp.KMSKeyArn != nil {
		cr.Spec.ForProvider.KMSKeyARN = resp.KMSKeyArn
	} else {
		cr.Spec.ForProvider.KMSKeyARN = nil
	}
	if resp.LastModified != nil {
		cr.Status.AtProvider.LastModified = resp.LastModified
	} else {
		cr.Status.AtProvider.LastModified = nil
	}
	if resp.LastUpdateStatus != nil {
		cr.Status.AtProvider.LastUpdateStatus = resp.LastUpdateStatus
	} else {
		cr.Status.AtProvider.LastUpdateStatus = nil
	}
	if resp.LastUpdateStatusReason != nil {
		cr.Status.AtProvider.LastUpdateStatusReason = resp.LastUpdateStatusReason
	} else {
		cr.Status.AtProvider.LastUpdateStatusReason = nil
	}
	if resp.LastUpdateStatusReasonCode != nil {
		cr.Status.AtProvider.LastUpdateStatusReasonCode = resp.LastUpdateStatusReasonCode
	} else {
		cr.Status.AtProvider.LastUpdateStatusReasonCode = nil
	}
	if resp.MasterArn != nil {
		cr.Status.AtProvider.MasterARN = resp.MasterArn
	} else {
		cr.Status.AtProvider.MasterARN = nil
	}
	if resp.MemorySize != nil {
		cr.Spec.ForProvider.MemorySize = resp.MemorySize
	} else {
		cr.Spec.ForProvider.MemorySize = nil
	}
	if resp.PackageType != nil {
		cr.Spec.ForProvider.PackageType = resp.PackageType
	} else {
		cr.Spec.ForProvider.PackageType = nil
	}
	if resp.RevisionId != nil {
		cr.Status.AtProvider.RevisionID = resp.RevisionId
	} else {
		cr.Status.AtProvider.RevisionID = nil
	}
	if resp.Role != nil {
		cr.Status.AtProvider.Role = resp.Role
	} else {
		cr.Status.AtProvider.Role = nil
	}
	if resp.Runtime != nil {
		cr.Spec.ForProvider.Runtime = resp.Runtime
	} else {
		cr.Spec.ForProvider.Runtime = nil
	}
	if resp.SigningJobArn != nil {
		cr.Status.AtProvider.SigningJobARN = resp.SigningJobArn
	} else {
		cr.Status.AtProvider.SigningJobARN = nil
	}
	if resp.SigningProfileVersionArn != nil {
		cr.Status.AtProvider.SigningProfileVersionARN = resp.SigningProfileVersionArn
	} else {
		cr.Status.AtProvider.SigningProfileVersionARN = nil
	}
	if resp.State != nil {
		cr.Status.AtProvider.State = resp.State
	} else {
		cr.Status.AtProvider.State = nil
	}
	if resp.StateReason != nil {
		cr.Status.AtProvider.StateReason = resp.StateReason
	} else {
		cr.Status.AtProvider.StateReason = nil
	}
	if resp.StateReasonCode != nil {
		cr.Status.AtProvider.StateReasonCode = resp.StateReasonCode
	} else {
		cr.Status.AtProvider.StateReasonCode = nil
	}
	if resp.Timeout != nil {
		cr.Spec.ForProvider.Timeout = resp.Timeout
	} else {
		cr.Spec.ForProvider.Timeout = nil
	}
	if resp.TracingConfig != nil {
		f27 := &svcapitypes.TracingConfig{}
		if resp.TracingConfig.Mode != nil {
			f27.Mode = resp.TracingConfig.Mode
		}
		cr.Spec.ForProvider.TracingConfig = f27
	} else {
		cr.Spec.ForProvider.TracingConfig = nil
	}
	if resp.Version != nil {
		cr.Status.AtProvider.Version = resp.Version
	} else {
		cr.Status.AtProvider.Version = nil
	}
	if resp.VpcConfig != nil {
		f29 := &svcapitypes.VPCConfigResponse{}
		if resp.VpcConfig.SecurityGroupIds != nil {
			f29f0 := []*string{}
			for _, f29f0iter := range resp.VpcConfig.SecurityGroupIds {
				var f29f0elem string
				f29f0elem = *f29f0iter
				f29f0 = append(f29f0, &f29f0elem)
			}
			f29.SecurityGroupIDs = f29f0
		}
		if resp.VpcConfig.SubnetIds != nil {
			f29f1 := []*string{}
			for _, f29f1iter := range resp.VpcConfig.SubnetIds {
				var f29f1elem string
				f29f1elem = *f29f1iter
				f29f1 = append(f29f1, &f29f1elem)
			}
			f29.SubnetIDs = f29f1
		}
		if resp.VpcConfig.VpcId != nil {
			f29.VPCID = resp.VpcConfig.VpcId
		}
		cr.Status.AtProvider.VPCConfig = f29
	} else {
		cr.Status.AtProvider.VPCConfig = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	return e.update(ctx, mg)

}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Function)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteFunctionInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteFunctionWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.LambdaAPI, opts []option) *external {
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
		update:         nopUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.LambdaAPI
	preObserve     func(context.Context, *svcapitypes.Function, *svcsdk.GetFunctionInput) error
	postObserve    func(context.Context, *svcapitypes.Function, *svcsdk.GetFunctionOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.FunctionParameters, *svcsdk.GetFunctionOutput) error
	isUpToDate     func(*svcapitypes.Function, *svcsdk.GetFunctionOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.Function, *svcsdk.CreateFunctionInput) error
	postCreate     func(context.Context, *svcapitypes.Function, *svcsdk.FunctionConfiguration, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Function, *svcsdk.DeleteFunctionInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Function, *svcsdk.DeleteFunctionOutput, error) error
	update         func(context.Context, cpresource.Managed) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Function, *svcsdk.GetFunctionInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Function, _ *svcsdk.GetFunctionOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.FunctionParameters, *svcsdk.GetFunctionOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.Function, *svcsdk.GetFunctionOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.Function, *svcsdk.CreateFunctionInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Function, _ *svcsdk.FunctionConfiguration, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Function, *svcsdk.DeleteFunctionInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Function, _ *svcsdk.DeleteFunctionOutput, err error) error {
	return err
}
func nopUpdate(context.Context, cpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}
