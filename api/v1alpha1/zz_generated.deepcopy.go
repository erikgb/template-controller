//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedResourceInfo) DeepCopyInto(out *AppliedResourceInfo) {
	*out = *in
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedResourceInfo.
func (in *AppliedResourceInfo) DeepCopy() *AppliedResourceInfo {
	if in == nil {
		return nil
	}
	out := new(AppliedResourceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Generator) DeepCopyInto(out *Generator) {
	*out = *in
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(PullRequestGenerator)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Generator.
func (in *Generator) DeepCopy() *Generator {
	if in == nil {
		return nil
	}
	out := new(Generator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubProject) DeepCopyInto(out *GithubProject) {
	*out = *in
	if in.TokenRef != nil {
		in, out := &in.TokenRef, &out.TokenRef
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubProject.
func (in *GithubProject) DeepCopy() *GithubProject {
	if in == nil {
		return nil
	}
	out := new(GithubProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubPullRequest1) DeepCopyInto(out *GithubPullRequest1) {
	*out = *in
	in.GithubProject.DeepCopyInto(&out.GithubProject)
	if in.PullRequestId != nil {
		in, out := &in.PullRequestId, &out.PullRequestId
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubPullRequest1.
func (in *GithubPullRequest1) DeepCopy() *GithubPullRequest1 {
	if in == nil {
		return nil
	}
	out := new(GithubPullRequest1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitlabMergeRequest1) DeepCopyInto(out *GitlabMergeRequest1) {
	*out = *in
	in.GitlabProject.DeepCopyInto(&out.GitlabProject)
	if in.MergeRequestId != nil {
		in, out := &in.MergeRequestId, &out.MergeRequestId
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitlabMergeRequest1.
func (in *GitlabMergeRequest1) DeepCopy() *GitlabMergeRequest1 {
	if in == nil {
		return nil
	}
	out := new(GitlabMergeRequest1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitlabProject) DeepCopyInto(out *GitlabProject) {
	*out = *in
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(string)
		**out = **in
	}
	if in.TokenRef != nil {
		in, out := &in.TokenRef, &out.TokenRef
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitlabProject.
func (in *GitlabProject) DeepCopy() *GitlabProject {
	if in == nil {
		return nil
	}
	out := new(GitlabProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Handler) DeepCopyInto(out *Handler) {
	*out = *in
	if in.PullRequestComment != nil {
		in, out := &in.PullRequestComment, &out.PullRequestComment
		*out = new(PullRequestCommentReporter)
		(*in).DeepCopyInto(*out)
	}
	if in.PullRequestApprove != nil {
		in, out := &in.PullRequestApprove, &out.PullRequestApprove
		*out = new(PullRequestApproveReporter)
		(*in).DeepCopyInto(*out)
	}
	if in.PullRequestCommand != nil {
		in, out := &in.PullRequestCommand, &out.PullRequestCommand
		*out = new(PullRequestCommandHandler)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Handler.
func (in *Handler) DeepCopy() *Handler {
	if in == nil {
		return nil
	}
	out := new(Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandlerStatus) DeepCopyInto(out *HandlerStatus) {
	*out = *in
	if in.PullRequestComment != nil {
		in, out := &in.PullRequestComment, &out.PullRequestComment
		*out = new(PullRequestCommentReporterStatus)
		**out = **in
	}
	if in.PullRequestApprove != nil {
		in, out := &in.PullRequestApprove, &out.PullRequestApprove
		*out = new(PullRequestApproveReporterStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.PullRequestCommand != nil {
		in, out := &in.PullRequestCommand, &out.PullRequestCommand
		*out = new(PullRequestCommandHandlerStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandlerStatus.
func (in *HandlerStatus) DeepCopy() *HandlerStatus {
	if in == nil {
		return nil
	}
	out := new(HandlerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectHandler) DeepCopyInto(out *ObjectHandler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectHandler.
func (in *ObjectHandler) DeepCopy() *ObjectHandler {
	if in == nil {
		return nil
	}
	out := new(ObjectHandler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectHandler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectHandlerDefaultsSpec) DeepCopyInto(out *ObjectHandlerDefaultsSpec) {
	*out = *in
	if in.Gitlab != nil {
		in, out := &in.Gitlab, &out.Gitlab
		*out = new(GitlabMergeRequest1)
		(*in).DeepCopyInto(*out)
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(GithubPullRequest1)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectHandlerDefaultsSpec.
func (in *ObjectHandlerDefaultsSpec) DeepCopy() *ObjectHandlerDefaultsSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectHandlerDefaultsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectHandlerList) DeepCopyInto(out *ObjectHandlerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectHandler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectHandlerList.
func (in *ObjectHandlerList) DeepCopy() *ObjectHandlerList {
	if in == nil {
		return nil
	}
	out := new(ObjectHandlerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectHandlerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectHandlerSpec) DeepCopyInto(out *ObjectHandlerSpec) {
	*out = *in
	out.Interval = in.Interval
	out.ForObject = in.ForObject
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = new(ObjectHandlerDefaultsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Handlers != nil {
		in, out := &in.Handlers, &out.Handlers
		*out = make([]Handler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectHandlerSpec.
func (in *ObjectHandlerSpec) DeepCopy() *ObjectHandlerSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectHandlerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectHandlerStatus) DeepCopyInto(out *ObjectHandlerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HandlerStatus != nil {
		in, out := &in.HandlerStatus, &out.HandlerStatus
		*out = make([]*HandlerStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HandlerStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectHandlerStatus.
func (in *ObjectHandlerStatus) DeepCopy() *ObjectHandlerStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectHandlerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTemplate) DeepCopyInto(out *ObjectTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTemplate.
func (in *ObjectTemplate) DeepCopy() *ObjectTemplate {
	if in == nil {
		return nil
	}
	out := new(ObjectTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTemplateDefaultsSpec) DeepCopyInto(out *ObjectTemplateDefaultsSpec) {
	*out = *in
	if in.Gitlab != nil {
		in, out := &in.Gitlab, &out.Gitlab
		*out = new(GitlabProject)
		(*in).DeepCopyInto(*out)
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(GithubProject)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTemplateDefaultsSpec.
func (in *ObjectTemplateDefaultsSpec) DeepCopy() *ObjectTemplateDefaultsSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectTemplateDefaultsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTemplateList) DeepCopyInto(out *ObjectTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTemplateList.
func (in *ObjectTemplateList) DeepCopy() *ObjectTemplateList {
	if in == nil {
		return nil
	}
	out := new(ObjectTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTemplateSpec) DeepCopyInto(out *ObjectTemplateSpec) {
	*out = *in
	out.Interval = in.Interval
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = new(ObjectTemplateDefaultsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Generators != nil {
		in, out := &in.Generators, &out.Generators
		*out = make([]Generator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]Template, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTemplateSpec.
func (in *ObjectTemplateSpec) DeepCopy() *ObjectTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTemplateStatus) DeepCopyInto(out *ObjectTemplateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AppliedResources != nil {
		in, out := &in.AppliedResources, &out.AppliedResources
		*out = make([]AppliedResourceInfo, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTemplateStatus.
func (in *ObjectTemplateStatus) DeepCopy() *ObjectTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestApproveReporter) DeepCopyInto(out *PullRequestApproveReporter) {
	*out = *in
	if in.Gitlab != nil {
		in, out := &in.Gitlab, &out.Gitlab
		*out = new(GitlabMergeRequest1)
		(*in).DeepCopyInto(*out)
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(GithubPullRequest1)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestApproveReporter.
func (in *PullRequestApproveReporter) DeepCopy() *PullRequestApproveReporter {
	if in == nil {
		return nil
	}
	out := new(PullRequestApproveReporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestApproveReporterStatus) DeepCopyInto(out *PullRequestApproveReporterStatus) {
	*out = *in
	if in.Approved != nil {
		in, out := &in.Approved, &out.Approved
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestApproveReporterStatus.
func (in *PullRequestApproveReporterStatus) DeepCopy() *PullRequestApproveReporterStatus {
	if in == nil {
		return nil
	}
	out := new(PullRequestApproveReporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestCommandHandler) DeepCopyInto(out *PullRequestCommandHandler) {
	*out = *in
	if in.Gitlab != nil {
		in, out := &in.Gitlab, &out.Gitlab
		*out = new(GitlabMergeRequest1)
		(*in).DeepCopyInto(*out)
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(GithubPullRequest1)
		(*in).DeepCopyInto(*out)
	}
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]PullRequestCommandHandlerCommandSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestCommandHandler.
func (in *PullRequestCommandHandler) DeepCopy() *PullRequestCommandHandler {
	if in == nil {
		return nil
	}
	out := new(PullRequestCommandHandler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestCommandHandlerActionAnnotateSpec) DeepCopyInto(out *PullRequestCommandHandlerActionAnnotateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestCommandHandlerActionAnnotateSpec.
func (in *PullRequestCommandHandlerActionAnnotateSpec) DeepCopy() *PullRequestCommandHandlerActionAnnotateSpec {
	if in == nil {
		return nil
	}
	out := new(PullRequestCommandHandlerActionAnnotateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestCommandHandlerActionSpec) DeepCopyInto(out *PullRequestCommandHandlerActionSpec) {
	*out = *in
	if in.Annotate != nil {
		in, out := &in.Annotate, &out.Annotate
		*out = new(PullRequestCommandHandlerActionAnnotateSpec)
		**out = **in
	}
	if in.JsonPatch != nil {
		in, out := &in.JsonPatch, &out.JsonPatch
		*out = new([]runtime.RawExtension)
		if **in != nil {
			in, out := *in, *out
			*out = make([]runtime.RawExtension, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestCommandHandlerActionSpec.
func (in *PullRequestCommandHandlerActionSpec) DeepCopy() *PullRequestCommandHandlerActionSpec {
	if in == nil {
		return nil
	}
	out := new(PullRequestCommandHandlerActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestCommandHandlerCommandSpec) DeepCopyInto(out *PullRequestCommandHandlerCommandSpec) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]PullRequestCommandHandlerActionSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestCommandHandlerCommandSpec.
func (in *PullRequestCommandHandlerCommandSpec) DeepCopy() *PullRequestCommandHandlerCommandSpec {
	if in == nil {
		return nil
	}
	out := new(PullRequestCommandHandlerCommandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestCommandHandlerStatus) DeepCopyInto(out *PullRequestCommandHandlerStatus) {
	*out = *in
	if in.LastProcessedCommentTime != nil {
		in, out := &in.LastProcessedCommentTime, &out.LastProcessedCommentTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestCommandHandlerStatus.
func (in *PullRequestCommandHandlerStatus) DeepCopy() *PullRequestCommandHandlerStatus {
	if in == nil {
		return nil
	}
	out := new(PullRequestCommandHandlerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestCommentReporter) DeepCopyInto(out *PullRequestCommentReporter) {
	*out = *in
	if in.Gitlab != nil {
		in, out := &in.Gitlab, &out.Gitlab
		*out = new(GitlabMergeRequest1)
		(*in).DeepCopyInto(*out)
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(GithubPullRequest1)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestCommentReporter.
func (in *PullRequestCommentReporter) DeepCopy() *PullRequestCommentReporter {
	if in == nil {
		return nil
	}
	out := new(PullRequestCommentReporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestCommentReporterStatus) DeepCopyInto(out *PullRequestCommentReporterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestCommentReporterStatus.
func (in *PullRequestCommentReporterStatus) DeepCopy() *PullRequestCommentReporterStatus {
	if in == nil {
		return nil
	}
	out := new(PullRequestCommentReporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestGenerator) DeepCopyInto(out *PullRequestGenerator) {
	*out = *in
	if in.Gitlab != nil {
		in, out := &in.Gitlab, &out.Gitlab
		*out = new(PullRequestGeneratorGitlab)
		(*in).DeepCopyInto(*out)
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(PullRequestGeneratorGithub)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestGenerator.
func (in *PullRequestGenerator) DeepCopy() *PullRequestGenerator {
	if in == nil {
		return nil
	}
	out := new(PullRequestGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestGeneratorGithub) DeepCopyInto(out *PullRequestGeneratorGithub) {
	*out = *in
	in.GithubProject.DeepCopyInto(&out.GithubProject)
	if in.TargetBranch != nil {
		in, out := &in.TargetBranch, &out.TargetBranch
		*out = new(string)
		**out = **in
	}
	if in.SourceBranch != nil {
		in, out := &in.SourceBranch, &out.SourceBranch
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestGeneratorGithub.
func (in *PullRequestGeneratorGithub) DeepCopy() *PullRequestGeneratorGithub {
	if in == nil {
		return nil
	}
	out := new(PullRequestGeneratorGithub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestGeneratorGitlab) DeepCopyInto(out *PullRequestGeneratorGitlab) {
	*out = *in
	in.GitlabProject.DeepCopyInto(&out.GitlabProject)
	if in.TargetBranch != nil {
		in, out := &in.TargetBranch, &out.TargetBranch
		*out = new(string)
		**out = **in
	}
	if in.SourceBranch != nil {
		in, out := &in.SourceBranch, &out.SourceBranch
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestGeneratorGitlab.
func (in *PullRequestGeneratorGitlab) DeepCopy() *PullRequestGeneratorGitlab {
	if in == nil {
		return nil
	}
	out := new(PullRequestGeneratorGitlab)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryGithubPullRequests) DeepCopyInto(out *QueryGithubPullRequests) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryGithubPullRequests.
func (in *QueryGithubPullRequests) DeepCopy() *QueryGithubPullRequests {
	if in == nil {
		return nil
	}
	out := new(QueryGithubPullRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryGithubPullRequests) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryGithubPullRequestsList) DeepCopyInto(out *QueryGithubPullRequestsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QueryGithubPullRequests, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryGithubPullRequestsList.
func (in *QueryGithubPullRequestsList) DeepCopy() *QueryGithubPullRequestsList {
	if in == nil {
		return nil
	}
	out := new(QueryGithubPullRequestsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryGithubPullRequestsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryGithubPullRequestsSpec) DeepCopyInto(out *QueryGithubPullRequestsSpec) {
	*out = *in
	out.Interval = in.Interval
	in.GithubProject.DeepCopyInto(&out.GithubProject)
	if in.Head != nil {
		in, out := &in.Head, &out.Head
		*out = new(string)
		**out = **in
	}
	if in.Base != nil {
		in, out := &in.Base, &out.Base
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryGithubPullRequestsSpec.
func (in *QueryGithubPullRequestsSpec) DeepCopy() *QueryGithubPullRequestsSpec {
	if in == nil {
		return nil
	}
	out := new(QueryGithubPullRequestsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryGithubPullRequestsStatus) DeepCopyInto(out *QueryGithubPullRequestsStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PullRequests != nil {
		in, out := &in.PullRequests, &out.PullRequests
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryGithubPullRequestsStatus.
func (in *QueryGithubPullRequestsStatus) DeepCopy() *QueryGithubPullRequestsStatus {
	if in == nil {
		return nil
	}
	out := new(QueryGithubPullRequestsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryGitlabMergeRequests) DeepCopyInto(out *QueryGitlabMergeRequests) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryGitlabMergeRequests.
func (in *QueryGitlabMergeRequests) DeepCopy() *QueryGitlabMergeRequests {
	if in == nil {
		return nil
	}
	out := new(QueryGitlabMergeRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryGitlabMergeRequests) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryGitlabMergeRequestsList) DeepCopyInto(out *QueryGitlabMergeRequestsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QueryGitlabMergeRequests, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryGitlabMergeRequestsList.
func (in *QueryGitlabMergeRequestsList) DeepCopy() *QueryGitlabMergeRequestsList {
	if in == nil {
		return nil
	}
	out := new(QueryGitlabMergeRequestsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryGitlabMergeRequestsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryGitlabMergeRequestsSpec) DeepCopyInto(out *QueryGitlabMergeRequestsSpec) {
	*out = *in
	out.Interval = in.Interval
	in.GitlabProject.DeepCopyInto(&out.GitlabProject)
	if in.TargetBranch != nil {
		in, out := &in.TargetBranch, &out.TargetBranch
		*out = new(string)
		**out = **in
	}
	if in.SourceBranch != nil {
		in, out := &in.SourceBranch, &out.SourceBranch
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryGitlabMergeRequestsSpec.
func (in *QueryGitlabMergeRequestsSpec) DeepCopy() *QueryGitlabMergeRequestsSpec {
	if in == nil {
		return nil
	}
	out := new(QueryGitlabMergeRequestsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryGitlabMergeRequestsStatus) DeepCopyInto(out *QueryGitlabMergeRequestsStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MergeRequests != nil {
		in, out := &in.MergeRequests, &out.MergeRequests
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryGitlabMergeRequestsStatus.
func (in *QueryGitlabMergeRequestsStatus) DeepCopy() *QueryGitlabMergeRequestsStatus {
	if in == nil {
		return nil
	}
	out := new(QueryGitlabMergeRequestsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = (*in).DeepCopy()
	}
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}
