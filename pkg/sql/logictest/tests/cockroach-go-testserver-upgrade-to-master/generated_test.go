// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by generate-logictest, DO NOT EDIT.

package testcockroach_go_testserver_upgrade_to_master

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/build/bazel"
	"github.com/cockroachdb/cockroach/pkg/security/securityassets"
	"github.com/cockroachdb/cockroach/pkg/security/securitytest"
	"github.com/cockroachdb/cockroach/pkg/server"
	"github.com/cockroachdb/cockroach/pkg/sql/logictest"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/randutil"
)

const configIdx = 18

var logicTestDir string

func init() {
	if bazel.BuiltWithBazel() {
		var err error
		logicTestDir, err = bazel.Runfile("pkg/sql/logictest/testdata/logic_test")
		if err != nil {
			panic(err)
		}
	} else {
		logicTestDir = "../../../../sql/logictest/testdata/logic_test"
	}
}

func TestMain(m *testing.M) {
	securityassets.SetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)

	defer serverutils.TestingSetDefaultTenantSelectionOverride(
		base.TestIsForStuffThatShouldWorkWithSecondaryTenantsButDoesntYet(76378),
	)()

	os.Exit(m.Run())
}

func runLogicTest(t *testing.T, file string) {
	skip.UnderDeadlock(t, "times out and/or hangs")
	logictest.RunLogicTest(t, logictest.TestServerArgs{}, configIdx, filepath.Join(logicTestDir, file))
}

// TestLogic_tmp runs any tests that are prefixed with "_", in which a dedicated
// test is not generated for. This allows developers to create and run temporary
// test files that are not checked into the repository, without repeatedly
// regenerating and reverting changes to this file, generated_test.go.
//
// TODO(mgartner): Add file filtering so that individual files can be run,
// instead of all files with the "_" prefix.
func TestLogic_tmp(t *testing.T) {
	defer leaktest.AfterTest(t)()
	var glob string
	glob = filepath.Join(logicTestDir, "_*")
	logictest.RunLogicTests(t, logictest.TestServerArgs{}, configIdx, glob)
}

func TestLogic_mixed_version_can_login(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_can_login")
}

func TestLogic_mixed_version_database_role_settings_role_id(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_database_role_settings_role_id")
}

func TestLogic_mixed_version_external_connections_owner_id(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_external_connections_owner_id")
}

func TestLogic_mixed_version_insights_queries(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_insights_queries")
}

func TestLogic_mixed_version_new_system_privileges(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_new_system_privileges")
}

func TestLogic_mixed_version_partially_visible_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_partially_visible_index")
}

func TestLogic_mixed_version_range_tombstones(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_range_tombstones")
}

func TestLogic_mixed_version_role_members_user_ids(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_role_members_user_ids")
}

func TestLogic_mixed_version_system_privileges_user_id(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runLogicTest(t, "mixed_version_system_privileges_user_id")
}
