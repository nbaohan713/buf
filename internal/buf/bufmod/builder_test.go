// Copyright 2020 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufmod

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/bufbuild/buf/internal/buf/bufcore"
	"github.com/bufbuild/buf/internal/buf/bufcore/bufcoretesting"
	"github.com/bufbuild/buf/internal/pkg/normalpath"
	"github.com/bufbuild/buf/internal/pkg/storage/storageos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestGetFileInfos1(t *testing.T) {
	testGetFileInfos(
		t,
		"testdata/1",
		[]string{
			"proto",
		},
		[]string{
			"proto/b",
		},
		bufcoretesting.NewFileInfo(t, "a/1.proto", "testdata/1/proto/a/1.proto", false),
		bufcoretesting.NewFileInfo(t, "a/2.proto", "testdata/1/proto/a/2.proto", false),
		bufcoretesting.NewFileInfo(t, "a/3.proto", "testdata/1/proto/a/3.proto", false),
		bufcoretesting.NewFileInfo(t, "a/c/1.proto", "testdata/1/proto/a/c/1.proto", false),
		bufcoretesting.NewFileInfo(t, "a/c/2.proto", "testdata/1/proto/a/c/2.proto", false),
		bufcoretesting.NewFileInfo(t, "a/c/3.proto", "testdata/1/proto/a/c/3.proto", false),
		bufcoretesting.NewFileInfo(t, "d/1.proto", "testdata/1/proto/d/1.proto", false),
		bufcoretesting.NewFileInfo(t, "d/2.proto", "testdata/1/proto/d/2.proto", false),
		bufcoretesting.NewFileInfo(t, "d/3.proto", "testdata/1/proto/d/3.proto", false),
	)
}

func TestGetFileInfos2(t *testing.T) {
	testGetFileInfos(
		t,
		"testdata/1",
		[]string{
			"proto",
		},
		[]string{
			"proto/b",
		},
		bufcoretesting.NewFileInfo(t, "a/1.proto", "testdata/1/proto/a/1.proto", false),
		bufcoretesting.NewFileInfo(t, "a/2.proto", "testdata/1/proto/a/2.proto", false),
		bufcoretesting.NewFileInfo(t, "a/3.proto", "testdata/1/proto/a/3.proto", false),
		bufcoretesting.NewFileInfo(t, "a/c/1.proto", "testdata/1/proto/a/c/1.proto", false),
		bufcoretesting.NewFileInfo(t, "a/c/2.proto", "testdata/1/proto/a/c/2.proto", false),
		bufcoretesting.NewFileInfo(t, "a/c/3.proto", "testdata/1/proto/a/c/3.proto", false),
		bufcoretesting.NewFileInfo(t, "d/1.proto", "testdata/1/proto/d/1.proto", false),
		bufcoretesting.NewFileInfo(t, "d/2.proto", "testdata/1/proto/d/2.proto", false),
		bufcoretesting.NewFileInfo(t, "d/3.proto", "testdata/1/proto/d/3.proto", false),
	)
}

func TestGetFileInfos3(t *testing.T) {
	testGetFileInfos(
		t,
		"testdata/1",
		[]string{
			"proto",
		},
		[]string{
			"proto/a",
		},
		bufcoretesting.NewFileInfo(t, "b/1.proto", "testdata/1/proto/b/1.proto", false),
		bufcoretesting.NewFileInfo(t, "b/2.proto", "testdata/1/proto/b/2.proto", false),
		bufcoretesting.NewFileInfo(t, "b/3.proto", "testdata/1/proto/b/3.proto", false),
		bufcoretesting.NewFileInfo(t, "d/1.proto", "testdata/1/proto/d/1.proto", false),
		bufcoretesting.NewFileInfo(t, "d/2.proto", "testdata/1/proto/d/2.proto", false),
		bufcoretesting.NewFileInfo(t, "d/3.proto", "testdata/1/proto/d/3.proto", false),
	)
}

func TestGetFileInfos4(t *testing.T) {
	testGetFileInfos(
		t,
		"testdata/1",
		[]string{
			"proto",
		},
		[]string{
			"proto/a/c",
		},
		bufcoretesting.NewFileInfo(t, "a/1.proto", "testdata/1/proto/a/1.proto", false),
		bufcoretesting.NewFileInfo(t, "a/2.proto", "testdata/1/proto/a/2.proto", false),
		bufcoretesting.NewFileInfo(t, "a/3.proto", "testdata/1/proto/a/3.proto", false),
		bufcoretesting.NewFileInfo(t, "b/1.proto", "testdata/1/proto/b/1.proto", false),
		bufcoretesting.NewFileInfo(t, "b/2.proto", "testdata/1/proto/b/2.proto", false),
		bufcoretesting.NewFileInfo(t, "b/3.proto", "testdata/1/proto/b/3.proto", false),
		bufcoretesting.NewFileInfo(t, "d/1.proto", "testdata/1/proto/d/1.proto", false),
		bufcoretesting.NewFileInfo(t, "d/2.proto", "testdata/1/proto/d/2.proto", false),
		bufcoretesting.NewFileInfo(t, "d/3.proto", "testdata/1/proto/d/3.proto", false),
	)
}

func TestGetFileInfos5(t *testing.T) {
	testGetFileInfos(
		t,
		"testdata/1",
		[]string{
			"proto",
		},
		[]string{
			"proto/a/c",
			"proto/d",
		},
		bufcoretesting.NewFileInfo(t, "a/1.proto", "testdata/1/proto/a/1.proto", false),
		bufcoretesting.NewFileInfo(t, "a/2.proto", "testdata/1/proto/a/2.proto", false),
		bufcoretesting.NewFileInfo(t, "a/3.proto", "testdata/1/proto/a/3.proto", false),
		bufcoretesting.NewFileInfo(t, "b/1.proto", "testdata/1/proto/b/1.proto", false),
		bufcoretesting.NewFileInfo(t, "b/2.proto", "testdata/1/proto/b/2.proto", false),
		bufcoretesting.NewFileInfo(t, "b/3.proto", "testdata/1/proto/b/3.proto", false),
	)
}

func TestGetAllFileInfosError1(t *testing.T) {
	testGetAllFileInfosError(
		t,
		"testdata/2",
		[]string{
			"a",
			"b",
		},
		[]string{},
	)
}

func TestGetFileInfosForExternalPaths1(t *testing.T) {
	testGetFileInfosForExternalPathsError(
		t,
		"testdata/2",
		[]string{
			"a",
			"b",
		},
		[]string{
			"testdata/2/a/1.proto",
			"testdata/2/a/2.proto",
			"testdata/2/a/3.proto",
			"testdata/2/b/1.proto",
			"testdata/2/b/4.proto",
		},
	)
}

func testGetFileInfos(
	t *testing.T,
	relDir string,
	relRoots []string,
	relExcludes []string,
	expectedFileInfos ...bufcore.FileInfo,
) {
	t.Parallel()
	readWriteBucket, err := storageos.NewReadWriteBucket(relDir)
	require.NoError(t, err)
	config, err := NewConfig(
		ExternalConfig{
			Roots:    relRoots,
			Excludes: relExcludes,
		},
	)
	require.NoError(t, err)
	module, err := NewBuilder(zap.NewNop()).BuildForBucket(
		context.Background(),
		readWriteBucket,
		config,
	)
	require.NoError(t, err)
	fileInfos, err := module.TargetFileInfos(context.Background())
	assert.NoError(t, err)
	bufcoretesting.AssertFileInfosEqual(
		t,
		expectedFileInfos,
		fileInfos,
	)
	if len(expectedFileInfos) > 1 {
		expectedFileInfos = expectedFileInfos[:len(expectedFileInfos)-1]
		bucketRelPaths := make([]string, len(expectedFileInfos))
		for i, expectedFileInfo := range expectedFileInfos {
			bucketRelExternalPath, err := filepath.Rel(relDir, expectedFileInfo.ExternalPath())
			require.NoError(t, err)
			bucketRelPath, err := normalpath.NormalizeAndValidate(bucketRelExternalPath)
			require.NoError(t, err)
			bucketRelPaths[i] = bucketRelPath
		}
		module, err := NewBuilder(zap.NewNop()).BuildForBucket(
			context.Background(),
			readWriteBucket,
			config,
			WithBucketRelPaths(bucketRelPaths...),
		)
		require.NoError(t, err)
		fileInfos, err := module.TargetFileInfos(context.Background())
		assert.NoError(t, err)
		bufcoretesting.AssertFileInfosEqual(
			t,
			expectedFileInfos,
			fileInfos,
		)
	}
}

func testGetAllFileInfosError(
	t *testing.T,
	relDir string,
	relRoots []string,
	relExcludes []string,
) {
	readWriteBucket, err := storageos.NewReadWriteBucket(relDir)
	require.NoError(t, err)
	config, err := NewConfig(
		ExternalConfig{
			Roots:    relRoots,
			Excludes: relExcludes,
		},
	)
	require.NoError(t, err)
	module, err := NewBuilder(zap.NewNop()).BuildForBucket(
		context.Background(),
		readWriteBucket,
		config,
	)
	require.NoError(t, err)
	_, err = module.TargetFileInfos(context.Background())
	assert.Error(t, err)
}

func testGetFileInfosForExternalPathsError(
	t *testing.T,
	relDir string,
	relRoots []string,
	externalPaths []string,
) {
	readWriteBucket, err := storageos.NewReadWriteBucket(relDir)
	require.NoError(t, err)
	config, err := NewConfig(
		ExternalConfig{
			Roots: relRoots,
		},
	)
	require.NoError(t, err)
	bucketRelPaths := make([]string, len(externalPaths))
	for i, externalPath := range externalPaths {
		bucketRelExternalPath, err := filepath.Rel(relDir, externalPath)
		require.NoError(t, err)
		bucketRelPath, err := normalpath.NormalizeAndValidate(bucketRelExternalPath)
		require.NoError(t, err)
		bucketRelPaths[i] = bucketRelPath
	}
	module, err := NewBuilder(zap.NewNop()).BuildForBucket(
		context.Background(),
		readWriteBucket,
		config,
		WithBucketRelPaths(bucketRelPaths...),
	)
	require.NoError(t, err)
	_, err = module.TargetFileInfos(context.Background())
	assert.Error(t, err)
}
