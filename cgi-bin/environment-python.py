#!/usr/bin/env python3

import os
import json

env_vars = dict(os.environ)

print("Content-Type: application/json\n")  # JSON output

print(json.dumps(env_vars, indent=2))
