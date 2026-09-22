from repositories.brand import addBrand


async def brand(data, db):
    return await addBrand(data, db)