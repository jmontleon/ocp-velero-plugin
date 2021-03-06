package main

import (
	"github.com/fusor/ocp-velero-plugin/velero-plugins/build"
	"github.com/fusor/ocp-velero-plugin/velero-plugins/common"
	"github.com/fusor/ocp-velero-plugin/velero-plugins/imagestream"
	"github.com/fusor/ocp-velero-plugin/velero-plugins/imagestreamtag"
	"github.com/fusor/ocp-velero-plugin/velero-plugins/pv"
	"github.com/fusor/ocp-velero-plugin/velero-plugins/route"
	veleroplugin "github.com/heptio/velero/pkg/plugin/framework"
	"github.com/sirupsen/logrus"
)

func main() {
	veleroplugin.NewServer().
		RegisterBackupItemAction("01-common-backup-plugin", newCommonBackupPlugin).
		RegisterRestoreItemAction("01-common-restore-plugin", newCommonRestorePlugin).
		RegisterBackupItemAction("02-pv-backup-plugin", newPVBackupPlugin).
		RegisterBackupItemAction("03-is-backup-plugin", newImageStreamBackupPlugin).
		RegisterRestoreItemAction("03-is-restore-plugin", newImageStreamRestorePlugin).
		RegisterRestoreItemAction("04-imagestreamtag-restore-plugin", newImageStreamTagRestorePlugin).
		RegisterRestoreItemAction("05-route-restore-plugin", newRouteRestorePlugin).
		RegisterRestoreItemAction("06-build-restore-plugin", newBuildRestorePlugin).
		Serve()
}

func newImageStreamBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &imagestream.BackupPlugin{Log: logger}, nil
}

func newImageStreamRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &imagestream.RestorePlugin{Log: logger}, nil
}

func newImageStreamTagRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &imagestreamtag.RestorePlugin{Log: logger}, nil
}

func newBuildRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &build.RestorePlugin{Log: logger}, nil
}

func newRouteRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &route.RestorePlugin{Log: logger}, nil
}

func newCommonBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &common.BackupPlugin{Log: logger}, nil
}

func newCommonRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &common.RestorePlugin{Log: logger}, nil
}

func newPVBackupPlugin(logger logrus.FieldLogger) (interface{}, error) {
	return &pv.BackupPlugin{Log: logger}, nil
}
