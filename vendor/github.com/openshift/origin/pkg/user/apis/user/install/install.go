package install

import (
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	kapi "k8s.io/kubernetes/pkg/api"

	"github.com/openshift/origin/pkg/api/legacy"
	userapi "github.com/openshift/origin/pkg/user/apis/user"
	userapiv1 "github.com/openshift/origin/pkg/user/apis/user/v1"
)

func init() {
	legacy.InstallLegacy(userapi.GroupName, userapi.AddToSchemeInCoreGroup, userapiv1.AddToSchemeInCoreGroup,
		sets.NewString("User", "Identity", "UserIdentityMapping", "Group"),
		kapi.Registry, kapi.Scheme,
	)
	Install(kapi.GroupFactoryRegistry, kapi.Registry, kapi.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  userapi.GroupName,
			VersionPreferenceOrder:     []string{userapiv1.SchemeGroupVersion.Version},
			AddInternalObjectsToScheme: userapi.AddToScheme,
			RootScopedKinds:            sets.NewString("User", "Identity", "UserIdentityMapping", "Group"),
		},
		announced.VersionToSchemeFunc{
			userapiv1.SchemeGroupVersion.Version: userapiv1.AddToScheme,
		},
	).Announce(groupFactoryRegistry).RegisterAndEnable(registry, scheme); err != nil {
		panic(err)
	}
}
