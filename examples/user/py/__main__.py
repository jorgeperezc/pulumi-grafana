"""A Python Pulumi program"""

import pulumi
import pulumi_grafana as grafana

user = grafana.User("test")

pulumi.export("user-name", user.name)
