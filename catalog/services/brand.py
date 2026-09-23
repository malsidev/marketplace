from repositories.brand import addBrand, brand


async def postBrand(data, db):
    return await addBrand(data, db)

async def getBrand(db):
    return await brand(db)