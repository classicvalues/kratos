/*

Copyright 2020 Adobe
All Rights Reserved.

NOTICE: Adobe permits you to use, modify, and distribute this file in
accordance with the terms of the Adobe license agreement accompanying
it. If you have received this file from a source other than Adobe,
then your use, modification, or distribution of it requires the prior
written permission of Adobe.

*/

package cache

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCache(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Cache Suite")
}

var _ = BeforeSuite(func() {

}, 60)

var _ = AfterSuite(func() {

})
