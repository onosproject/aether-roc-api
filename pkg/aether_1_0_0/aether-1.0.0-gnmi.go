// Code generated by oapi-codegen. DO NOT EDIT.
// Package aether_1_0_0 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package aether_1_0_0

import (
	"context"
)

import (
	"github.com/onosproject/aether-roc-api/pkg/gnmiutils"
)

// gnmiDeleteAetherV100targetAccessProfile deletes an instance of AetherV100targetAccessProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetAccessProfile, error) {
	var response AetherV100targetAccessProfile

	return &response, nil
}

// gnmiGetAetherV100targetAccessProfile returns an instance of AetherV100targetAccessProfile.
func (i *ServerImpl) gnmiGetAetherV100targetAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetAccessProfile, error) {
	var response AetherV100targetAccessProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetAccessProfile adds an instance of AetherV100targetAccessProfile.
func (w *ServerImpl) gnmiPostAetherV100targetAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetAccessProfile, error) {
	var response AetherV100targetAccessProfile

	return &response, nil
}

// gnmiDeleteAetherV100targetAccessProfileAccessProfile deletes an instance of AetherV100targetAccessProfileAccessProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetAccessProfileAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetAccessProfileAccessProfile, error) {
	var response AetherV100targetAccessProfileAccessProfile

	return &response, nil
}

// gnmiGetAetherV100targetAccessProfileAccessProfile returns an instance of AetherV100targetAccessProfileAccessProfile.
func (i *ServerImpl) gnmiGetAetherV100targetAccessProfileAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetAccessProfileAccessProfile, error) {
	var response AetherV100targetAccessProfileAccessProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetAccessProfileAccessProfile adds an instance of AetherV100targetAccessProfileAccessProfile.
func (w *ServerImpl) gnmiPostAetherV100targetAccessProfileAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetAccessProfileAccessProfile, error) {
	var response AetherV100targetAccessProfileAccessProfile

	return &response, nil
}

// gnmiDeleteAetherV100targetApnProfile deletes an instance of AetherV100targetApnProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetApnProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetApnProfile, error) {
	var response AetherV100targetApnProfile

	return &response, nil
}

// gnmiGetAetherV100targetApnProfile returns an instance of AetherV100targetApnProfile.
func (i *ServerImpl) gnmiGetAetherV100targetApnProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetApnProfile, error) {
	var response AetherV100targetApnProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetApnProfile adds an instance of AetherV100targetApnProfile.
func (w *ServerImpl) gnmiPostAetherV100targetApnProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetApnProfile, error) {
	var response AetherV100targetApnProfile

	return &response, nil
}

// gnmiDeleteAetherV100targetApnProfileApnProfile deletes an instance of AetherV100targetApnProfileApnProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetApnProfileApnProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetApnProfileApnProfile, error) {
	var response AetherV100targetApnProfileApnProfile

	return &response, nil
}

// gnmiGetAetherV100targetApnProfileApnProfile returns an instance of AetherV100targetApnProfileApnProfile.
func (i *ServerImpl) gnmiGetAetherV100targetApnProfileApnProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetApnProfileApnProfile, error) {
	var response AetherV100targetApnProfileApnProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetApnProfileApnProfile adds an instance of AetherV100targetApnProfileApnProfile.
func (w *ServerImpl) gnmiPostAetherV100targetApnProfileApnProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetApnProfileApnProfile, error) {
	var response AetherV100targetApnProfileApnProfile

	return &response, nil
}

// gnmiDeleteAetherV100targetQosProfile deletes an instance of AetherV100targetQosProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetQosProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfile, error) {
	var response AetherV100targetQosProfile

	return &response, nil
}

// gnmiGetAetherV100targetQosProfile returns an instance of AetherV100targetQosProfile.
func (i *ServerImpl) gnmiGetAetherV100targetQosProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfile, error) {
	var response AetherV100targetQosProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetQosProfile adds an instance of AetherV100targetQosProfile.
func (w *ServerImpl) gnmiPostAetherV100targetQosProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfile, error) {
	var response AetherV100targetQosProfile

	return &response, nil
}

// gnmiDeleteAetherV100targetQosProfileQosProfile deletes an instance of AetherV100targetQosProfileQosProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetQosProfileQosProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfileQosProfile, error) {
	var response AetherV100targetQosProfileQosProfile

	return &response, nil
}

// gnmiGetAetherV100targetQosProfileQosProfile returns an instance of AetherV100targetQosProfileQosProfile.
func (i *ServerImpl) gnmiGetAetherV100targetQosProfileQosProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfileQosProfile, error) {
	var response AetherV100targetQosProfileQosProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetQosProfileQosProfile adds an instance of AetherV100targetQosProfileQosProfile.
func (w *ServerImpl) gnmiPostAetherV100targetQosProfileQosProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfileQosProfile, error) {
	var response AetherV100targetQosProfileQosProfile

	return &response, nil
}

// gnmiDeleteAetherV100targetQosProfileQosProfileApnAmbr deletes an instance of AetherV100targetQosProfileQosProfileApnAmbr.
func (w *ServerImpl) gnmiDeleteAetherV100targetQosProfileQosProfileApnAmbr(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfileQosProfileApnAmbr, error) {
	var response AetherV100targetQosProfileQosProfileApnAmbr

	return &response, nil
}

// gnmiGetAetherV100targetQosProfileQosProfileApnAmbr returns an instance of AetherV100targetQosProfileQosProfileApnAmbr.
func (i *ServerImpl) gnmiGetAetherV100targetQosProfileQosProfileApnAmbr(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfileQosProfileApnAmbr, error) {
	var response AetherV100targetQosProfileQosProfileApnAmbr

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetQosProfileQosProfileApnAmbr adds an instance of AetherV100targetQosProfileQosProfileApnAmbr.
func (w *ServerImpl) gnmiPostAetherV100targetQosProfileQosProfileApnAmbr(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetQosProfileQosProfileApnAmbr, error) {
	var response AetherV100targetQosProfileQosProfileApnAmbr

	return &response, nil
}

// gnmiDeleteAetherV100targetSubscriber deletes an instance of AetherV100targetSubscriber.
func (w *ServerImpl) gnmiDeleteAetherV100targetSubscriber(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriber, error) {
	var response AetherV100targetSubscriber

	return &response, nil
}

// gnmiGetAetherV100targetSubscriber returns an instance of AetherV100targetSubscriber.
func (i *ServerImpl) gnmiGetAetherV100targetSubscriber(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriber, error) {
	var response AetherV100targetSubscriber

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetSubscriber adds an instance of AetherV100targetSubscriber.
func (w *ServerImpl) gnmiPostAetherV100targetSubscriber(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriber, error) {
	var response AetherV100targetSubscriber

	return &response, nil
}

// gnmiDeleteAetherV100targetSubscriberUe deletes an instance of AetherV100targetSubscriberUe.
func (w *ServerImpl) gnmiDeleteAetherV100targetSubscriberUe(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUe, error) {
	var response AetherV100targetSubscriberUe

	return &response, nil
}

// gnmiGetAetherV100targetSubscriberUe returns an instance of AetherV100targetSubscriberUe.
func (i *ServerImpl) gnmiGetAetherV100targetSubscriberUe(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUe, error) {
	var response AetherV100targetSubscriberUe

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetSubscriberUe adds an instance of AetherV100targetSubscriberUe.
func (w *ServerImpl) gnmiPostAetherV100targetSubscriberUe(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUe, error) {
	var response AetherV100targetSubscriberUe

	return &response, nil
}

// gnmiDeleteAetherV100targetSubscriberUeProfiles deletes an instance of AetherV100targetSubscriberUeProfiles.
func (w *ServerImpl) gnmiDeleteAetherV100targetSubscriberUeProfiles(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeProfiles, error) {
	var response AetherV100targetSubscriberUeProfiles

	return &response, nil
}

// gnmiGetAetherV100targetSubscriberUeProfiles returns an instance of AetherV100targetSubscriberUeProfiles.
func (i *ServerImpl) gnmiGetAetherV100targetSubscriberUeProfiles(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeProfiles, error) {
	var response AetherV100targetSubscriberUeProfiles

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetSubscriberUeProfiles adds an instance of AetherV100targetSubscriberUeProfiles.
func (w *ServerImpl) gnmiPostAetherV100targetSubscriberUeProfiles(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeProfiles, error) {
	var response AetherV100targetSubscriberUeProfiles

	return &response, nil
}

// gnmiDeleteAetherV100targetSubscriberUeProfilesAccessProfile deletes an instance of AetherV100targetSubscriberUeProfilesAccessProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetSubscriberUeProfilesAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeProfilesAccessProfile, error) {
	var response AetherV100targetSubscriberUeProfilesAccessProfile

	return &response, nil
}

// gnmiGetAetherV100targetSubscriberUeProfilesAccessProfile returns an instance of AetherV100targetSubscriberUeProfilesAccessProfile.
func (i *ServerImpl) gnmiGetAetherV100targetSubscriberUeProfilesAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeProfilesAccessProfile, error) {
	var response AetherV100targetSubscriberUeProfilesAccessProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetSubscriberUeProfilesAccessProfile adds an instance of AetherV100targetSubscriberUeProfilesAccessProfile.
func (w *ServerImpl) gnmiPostAetherV100targetSubscriberUeProfilesAccessProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeProfilesAccessProfile, error) {
	var response AetherV100targetSubscriberUeProfilesAccessProfile

	return &response, nil
}

// gnmiDeleteAetherV100targetSubscriberUeServingPlmn deletes an instance of AetherV100targetSubscriberUeServingPlmn.
func (w *ServerImpl) gnmiDeleteAetherV100targetSubscriberUeServingPlmn(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeServingPlmn, error) {
	var response AetherV100targetSubscriberUeServingPlmn

	return &response, nil
}

// gnmiGetAetherV100targetSubscriberUeServingPlmn returns an instance of AetherV100targetSubscriberUeServingPlmn.
func (i *ServerImpl) gnmiGetAetherV100targetSubscriberUeServingPlmn(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeServingPlmn, error) {
	var response AetherV100targetSubscriberUeServingPlmn

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetSubscriberUeServingPlmn adds an instance of AetherV100targetSubscriberUeServingPlmn.
func (w *ServerImpl) gnmiPostAetherV100targetSubscriberUeServingPlmn(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetSubscriberUeServingPlmn, error) {
	var response AetherV100targetSubscriberUeServingPlmn

	return &response, nil
}

// gnmiDeleteAetherV100targetUpProfile deletes an instance of AetherV100targetUpProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetUpProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetUpProfile, error) {
	var response AetherV100targetUpProfile

	return &response, nil
}

// gnmiGetAetherV100targetUpProfile returns an instance of AetherV100targetUpProfile.
func (i *ServerImpl) gnmiGetAetherV100targetUpProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetUpProfile, error) {
	var response AetherV100targetUpProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetUpProfile adds an instance of AetherV100targetUpProfile.
func (w *ServerImpl) gnmiPostAetherV100targetUpProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetUpProfile, error) {
	var response AetherV100targetUpProfile

	return &response, nil
}

// gnmiDeleteAetherV100targetUpProfileUpProfile deletes an instance of AetherV100targetUpProfileUpProfile.
func (w *ServerImpl) gnmiDeleteAetherV100targetUpProfileUpProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetUpProfileUpProfile, error) {
	var response AetherV100targetUpProfileUpProfile

	return &response, nil
}

// gnmiGetAetherV100targetUpProfileUpProfile returns an instance of AetherV100targetUpProfileUpProfile.
func (i *ServerImpl) gnmiGetAetherV100targetUpProfileUpProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetUpProfileUpProfile, error) {
	var response AetherV100targetUpProfileUpProfile

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostAetherV100targetUpProfileUpProfile adds an instance of AetherV100targetUpProfileUpProfile.
func (w *ServerImpl) gnmiPostAetherV100targetUpProfileUpProfile(ctx context.Context, openApiPath string, target Target, args ...string) (*AetherV100targetUpProfileUpProfile, error) {
	var response AetherV100targetUpProfileUpProfile

	return &response, nil
}

// gnmiDeleteTarget deletes an instance of target.
func (w *ServerImpl) gnmiDeleteTarget(ctx context.Context, openApiPath string, target Target, args ...string) (*Target, error) {
	var response Target

	return &response, nil
}

// gnmiGetTarget returns an instance of target.
func (i *ServerImpl) gnmiGetTarget(ctx context.Context, openApiPath string, target Target, args ...string) (*Target, error) {
	var response Target

	gnmiGet, err := gnmiutils.NewGnmiGetRequest(openApiPath, string(target), args...)
	if err != nil {
		return nil, err
	}
	log.Infof("gnmiGetRequest %s", gnmiGet.String())
	update, err := gnmiutils.GetResponseUpdate(i.GnmiProvisioner.Get(ctx, gnmiGet))
	if err != nil {
		return nil, err
	}
	log.Info(update.String())

	return &response, nil
}

// gnmiPostTarget adds an instance of target.
func (w *ServerImpl) gnmiPostTarget(ctx context.Context, openApiPath string, target Target, args ...string) (*Target, error) {
	var response Target

	return &response, nil
}
