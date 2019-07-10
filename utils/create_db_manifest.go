package utils

import (
	"berty.tech/go-ipfs-log/io"
	"context"
	"github.com/berty/go-orbit-db/ipfs"
	"github.com/ipfs/go-cid"
	cbornode "github.com/ipfs/go-ipld-cbor"
	"github.com/pkg/errors"
	"github.com/polydawn/refmt/obj/atlas"
	"path"
)

type Manifest struct {
	Name             string
	Type             string
	AccessController string
}

func CreateDBManifest(ctx context.Context, services ipfs.Services, name string, dbType string, accessControllerAddress string, options interface{}) (cid.Cid, error) {
	manifest := &Manifest{
		Name:             name,
		Type:             dbType,
		AccessController: path.Join("/ipfs", accessControllerAddress),
	}

	c, err := io.WriteCBOR(ctx, services, manifest)
	if err != nil {
		return cid.Cid{}, errors.Wrap(err, "unable to write cbor data")
	}

	return c, err
}

var AtlasManifest = atlas.BuildEntry(Manifest{}).
	StructMap().
	AddField("Name", atlas.StructMapEntry{SerialName: "name"}).
	AddField("Type", atlas.StructMapEntry{SerialName: "type"}).
	AddField("AccessController", atlas.StructMapEntry{SerialName: "access_controller"}).
	Complete()

func init() {
	cbornode.RegisterCborType(AtlasManifest)
}
