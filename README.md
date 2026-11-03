# Gluster Cluster Manager (in Go)

This project is in very early stages and a learn as you go scenario for me (I am new to golang)

Most of this project's direction and process can be found in `planner.md`
An excerpt from that document:

**GOAL**: Gluster is intended to be a cluster/multi server/pc managment system for tasks and commands

Think k8s but designed for homelabs and their desired tasks in a fixed system domain, with linear execution.

Gluster is meant to be run on all devices that are meant to be in the cluster, with one main orchestrator gluster instance, then all the worker instances.
