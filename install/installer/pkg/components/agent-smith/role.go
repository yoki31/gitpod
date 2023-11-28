// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package agentsmith

import (
	"github.com/gitpod-io/gitpod/installer/pkg/common"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func role(ctx *common.RenderContext) ([]runtime.Object, error) {
	var rules []rbacv1.PolicyRule

	return []runtime.Object{
		&rbacv1.Role{
			TypeMeta: common.TypeMetaRole,
			ObjectMeta: metav1.ObjectMeta{
				Name:      Component,
				Namespace: ctx.Namespace,
				Labels:    common.DefaultLabels(Component),
			},
			Rules: append(rules, rbacv1.PolicyRule{
				APIGroups: []string{""},
				Resources: []string{"pods"},
				Verbs:     []string{"get", "update"},
			}),
		},
	}, nil
}
