from repositories.catalog import catalog


async def getCatalog(db):
    return await catalog(db)