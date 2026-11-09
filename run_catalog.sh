#!/usr/bin/env bash
cd /home/malsi/marketplace/catalog
source .venv/bin/activate
exec uvicorn main:app --host 0.0.0.0 --port 8001
