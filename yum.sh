#!/bin/bash

DIR=/etc/yum.repos.d

echo "$DIR/fedora.repo"
mv $DIR/fedora.repo $DIR/fedora.repo.bat
mv $DIR/fedora-updates.repo $DIR/fedora-updates.repo.bat
mv $DIR/fedora-modular.repo $DIR/fedora-modular.repo.bat
mv $DIR/fedora-updates-modular.repo $DIR/fedora-updates-modular.repo.bat

#cat > /etc/yum.repos.d/fedora.repo << EOF
#[fedora]
#name=Fedora $releasever - $basearch
#failovermethod=priority
#baseurl=https://mirrors.tuna.tsinghua.edu.cn/fedora/releases/$releasever/Everything/$basearch/os/
#metadata_expire=28d
#gpgcheck=1
#gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-fedora-$releasever-$basearch
#skip_if_unavailable=False
#EOF
#
#cat > /etc/yum.repos.d/fedora-updates.repo << EOF
#[updates]
#name=Fedora $releasever - $basearch - Updates
#failovermethod=priority
#baseurl=https://mirrors.tuna.tsinghua.edu.cn/fedora/updates/$releasever/Everything/$basearch/
#enabled=1
#gpgcheck=1
#metadata_expire=6h
#gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-fedora-$releasever-$basearch
#skip_if_unavailable=False
#EOF
#
#
#cat > /etc/yum.repos.d/fedora-modular.repo << EOF
#[fedora-modular]
#name=Fedora Modular $releasever - $basearch
#failovermethod=priority
#baseurl=https://mirrors.tuna.tsinghua.edu.cn/fedora/releases/$releasever/Modular/$basearch/os/
#enabled=1
#metadata_expire=7d
#gpgcheck=1
#gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-fedora-$releasever-$basearch
#skip_if_unavailable=False
#EOF
#
#
#cat > /etc/yum.repos.d/fedora-updates-modular.repo << EOF
#[updates-modular]
#name=Fedora Modular $releasever - $basearch - Updates
#failovermethod=priority
#baseurl=https://mirrors.tuna.tsinghua.edu.cn/fedora/updates/$releasever/Modular/$basearch/
#enabled=1
#gpgcheck=1
#metadata_expire=6h
#gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-fedora-$releasever-$basearch
#skip_if_unavailable=False
#EOF