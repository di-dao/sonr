SHELL=/bin/bash

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
SONR_ROOT_DIR=/Users/prad/Developer
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CORE_DIR=$(SONR_ROOT_DIR)/core
DESKTOP_DIR=$(SONR_ROOT_DIR)/desktop
MOBILE_DIR=$(SONR_ROOT_DIR)/mobile
CORE_FULL_DIR=$(SONR_ROOT_DIR)/core/cmd/daemon
CORE_BIND_DIR=$(SONR_ROOT_DIR)/core/cmd/bind
ELECTRON_BIN_DIR=$(SONR_ROOT_DIR)/electron/assets/bin/darwin

# Set this -->[/Users/xxxx/Sonr/]<-- to Folder of Sonr Repos
PROTO_DEF_PATH=/Users/prad/Developer/core/proto
APP_ROOT_DIR =/Users/prad/Developer/mobile/

# @ Packaging Vars/Commands
GOMOBILE=gomobile
GOCLEAN=$(GOMOBILE) clean
GOBIND=$(GOMOBILE) bind -ldflags='-s -w' -v
GOBIND_ANDROID=$(GOBIND) -target=android/arm64
GOBIND_IOS=$(GOBIND) -target=ios -bundleid=io.sonr.core

# @ Bind Directories
BIND_DIR_ANDROID=$(SONR_ROOT_DIR)/mobile/android/libs
BIND_DIR_IOS=$(SONR_ROOT_DIR)/mobile/ios/Frameworks
BIND_IOS_ARTIFACT= $(BIND_DIR_IOS)/Core.xcframework
BIND_ANDROID_ARTIFACT= $(BIND_DIR_ANDROID)/io.sonr.core.aar

# @ Proto Directories
PROTO_LIST_ALL=${ROOT_DIR}/proto/**/*.proto
MODULE_NAME=github.com/sonr-io/core
GO_OPT_FLAG=--go_opt=module=${MODULE_NAME}
GRPC_OPT_FLAG=--go-grpc_opt=module=${MODULE_NAME}
PROTO_GEN_GO="--go_out=."
PROTO_GEN_RPC="--go-grpc_out=."
PROTO_GEN_DOCS="--doc_out=docs"


# @ Distribution Release Variables
DIST_DIR=$(SONR_ROOT_DIR)/core/cmd/rpc/dist
DIST_DIR_DARWIN_AMD=$(DIST_DIR)/sonr-rpc_darwin_amd64
DIST_DIR_DARWIN_ARM=$(DIST_DIR)/sonr-rpc_darwin_arm64
DIST_DIR_LINUX_AMD=$(DIST_DIR)/sonr-rpc_linux_amd64
DIST_DIR_LINUX_ARM=$(DIST_DIR)/sonr-rpc_linux_arm64
DIST_DIR_WIN=$(DIST_DIR)/sonr-rpc_windows_amd64
DIST_ZIP_WIN=$(DIST_DIR)/*.zip

all: Makefile
	@figlet -f larry3d Sonr Core
	@echo ''
	@sed -n 's/^##//p ' $<

## bind        :   Binds Android and iOS for Plugin Path
bind: protobuf bind.ios bind.android
	@go mod tidy
	@cd /System/Library/Sounds && afplay Glass.aiff
	@echo ""
	@echo ""
	@echo "----------------------------------------------------------------"
	@echo "-------- ✅ ✅ ✅  SUCCESFUL MOBILE BIND  ✅ ✅ ✅  --------------"
	@echo "----------------------------------------------------------------"


## └─ android       - Android AAR
bind.android:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "--------------- 🤖 START ANDROID BIND 🤖 ----------------------"
	@echo "--------------------------------------------------------------"
	@go get golang.org/x/mobile/bind
	@gomobile init
	cd $(CORE_BIND_DIR) && $(GOBIND_ANDROID) -o $(BIND_ANDROID_ARTIFACT)
	@echo "✅ Finished Binding ➡ `date`"
	@echo ""


## └─ ios           - iOS Framework
bind.ios:
	@echo ""
	@echo ""
	@echo "--------------------------------------------------------------"
	@echo "-------------- 📱 START IOS BIND 📱 ---------------------------"
	@echo "--------------------------------------------------------------"
	@go get golang.org/x/mobile/bind
	cd $(CORE_BIND_DIR) && $(GOBIND_IOS) -o $(BIND_IOS_ARTIFACT)
	@echo "✅ Finished Binding ➡ `date`"
	@echo ""

dependencies:
	@echo '----------------------------------------'
	@echo '   Installing dependencies              '
	@echo '----------------------------------------'
	@GO111MODULE=off go get -u golang.org/x/mobile/cmd/...
	$(GOMOBILE_CMD) init -v

##
## [protobuf]     :   Compiles Protobuf models for Core Library and Plugin
protobuf:
	@echo "----"
	@echo "Sonr: Compiling Protobufs"
	@echo "----"
	@echo "Generating Protobuf Go code..."
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_GO) $(GO_OPT_FLAG)
	@protoc $(PROTO_LIST_ALL) --proto_path=$(ROOT_DIR) $(PROTO_GEN_RPC) $(GRPC_OPT_FLAG)

## [clean]     :   Reinitializes Gomobile and Removes Framworks from Plugin
clean:
	cd $(CORE_BIND_DIR) && $(GOCLEAN)
	go mod tidy
	go clean -cache -x
	rm -rf $(BIND_DIR_IOS)
	rm -rf $(BIND_DIR_ANDROID)
	mkdir -p $(BIND_DIR_IOS)
	mkdir -p $(BIND_DIR_ANDROID)
	cd $(CORE_BIND_DIR) && gomobile init
