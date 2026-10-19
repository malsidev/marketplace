from models.models import Products
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select

async def getCartProduct(product_id, db: AsyncSession):
    stmt = select(Products).where(Products.id.in_(product_id))
    data = await db.execute(stmt)
    res = data.scalars().all()
    return [
        {
            "id" : i.id,
            "name" : i.name,
            "description" : i.description,
        }
        for i in res
    ]