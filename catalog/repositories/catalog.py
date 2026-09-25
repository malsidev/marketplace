from sqlalchemy import select
from models.models import Brands, Categories, Products
from sqlalchemy.ext.asyncio import AsyncSession

async def catalog(db: AsyncSession):
    stmt = (
        select(Products, Brands, Categories)
        .join(Brands, Products.brand_id == Brands.id)
        .join(Categories, Products.category_id == Categories.id)
    )

    res = await db.execute(stmt)

    data = res.all()

    return [
        {
            "id": product.id,
            "name": product.name,
            "description": product.description,
            "brand": brand.name,
            "category": category.name,
        }
        for product, brand, category in data
    ]