from fastapi import FastAPI

from routers.catalog import router as catalog_router
from routers.products import router as products_router
from routers.brand import router as brand_router
from routers.categories import router as categories_router



def setup_routers(app: FastAPI):
    app.include_router(catalog_router)
    app.include_router(products_router)
    app.include_router(brand_router)
    app.include_router(categories_router)