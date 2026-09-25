from sqlalchemy import select
from models.models import Brands, Categories, Products
from sqlalchemy.ext.asyncio import AsyncSession
from database.redisdb import redis
import json
async def catalog(db: AsyncSession):

    cached = await redis.get("catalog")
    if cached:
        return json.loads(cached)
    stmt = (
        select(Products, Brands, Categories)
        .join(Brands, Products.brand_id == Brands.id)
        .join(Categories, Products.category_id == Categories.id)
    )

    res = await db.execute(stmt)

    rows = res.all()

    data = [
        {
            "id": product.id,
            "name": product.name,
            "description": product.description,
            "brand": brand.name,
            "category": category.name,
        }
        for product, brand, category in rows
    ]

    await redis.set(
        "catalog",
        json.dumps(data),
        ex=300
    )

    return data