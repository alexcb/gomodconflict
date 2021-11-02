# gomodconflict

Prints out conflicting version requirements between multiple go.mod files.

## Building

First download earthly.

Then run:

    earthly +all

builds are written to `build/ubuntu/gomodconflict` and `build/alpine/gomodconflict`

## Example use

    gomodconflict ~/gh/earthly/earthly/go.mod ~/earthly-hack/fsutil/go.mod ~/earthly-hack/buildkit/go.mod

which will display any conflicting go.mod values:

    github.com/containerd/continuity
      /home/alex/earthly-hack/fsutil/go.mod: v0.1.0
      /home/alex/earthly-hack/buildkit/go.mod: v0.2.0
    github.com/google/uuid
      /home/alex/gh/earthly/earthly/go.mod: v1.3.0
      /home/alex/earthly-hack/buildkit/go.mod: v1.2.0
    golang.org/x/sys
      /home/alex/earthly-hack/fsutil/go.mod: v0.0.0-20210313202042-bd2e13477e9c
      /home/alex/earthly-hack/buildkit/go.mod: v0.0.0-20210915083310-ed5796bab164
    github.com/docker/docker
      /home/alex/gh/earthly/earthly/go.mod: github.com/docker/docker=>v20.10.3-0.20210817025855-ba2adeebdb8d+incompatible
      /home/alex/earthly-hack/fsutil/go.mod: v20.10.3-0.20210817025855-ba2adeebdb8d+incompatible
      /home/alex/earthly-hack/buildkit/go.mod: github.com/docker/docker=>v20.10.3-0.20210817025855-ba2adeebdb8d+incompatible
    google.golang.org/protobuf
      /home/alex/gh/earthly/earthly/go.mod: v1.27.1
      /home/alex/earthly-hack/fsutil/go.mod: v1.25.0
      /home/alex/earthly-hack/buildkit/go.mod: v1.27.1
    github.com/containerd/stargz-snapshotter/estargz
      /home/alex/gh/earthly/earthly/go.mod: github.com/containerd/stargz-snapshotter/estargz=>v0.0.0-20201217071531-2b97b583765b
      /home/alex/earthly-hack/buildkit/go.mod: v0.8.1-0.20210910092506-a3ecdc9366fb
    github.com/golang/protobuf
      /home/alex/gh/earthly/earthly/go.mod: v1.5.2
      /home/alex/earthly-hack/fsutil/go.mod: v1.4.3
      /home/alex/earthly-hack/buildkit/go.mod: v1.5.2
    github.com/google/go-cmp
      /home/alex/earthly-hack/fsutil/go.mod: v0.5.2
      /home/alex/earthly-hack/buildkit/go.mod: v0.5.6

and will exit with an error code of 1 if differences are detected.
