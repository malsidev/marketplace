import { useEffect, useMemo, useState } from 'react'
import './App.css'

const API = {
  auth: 'http://localhost:8000',
  catalog: 'http://localhost:8001',
  orders: 'http://localhost:8080',
}

const fallbackProducts = [
  { id: 1, name: 'MacBook Pro 14', description: 'Ноутбук для работы и дизайна', brand: 'Apple', category: 'Tech', image: 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&w=900&q=80' },
  { id: 2, name: 'AirPods Pro', description: 'Беспроводные наушники с шумоподавлением', brand: 'Apple', category: 'Audio', image: 'https://images.unsplash.com/photo-1606220588913-b3aacb4d2f46?auto=format&fit=crop&w=900&q=80' },
  { id: 3, name: 'Bose QuietComfort', description: 'Комфортные наушники для поездок и дома', brand: 'Bose', category: 'Audio', image: 'https://images.unsplash.com/photo-1546435770-a3e426bf472b?auto=format&fit=crop&w=900&q=80' },
  { id: 4, name: 'iPhone 15', description: 'Флагманский смартфон со стабильной камерой', brand: 'Apple', category: 'Phones', image: 'https://images.unsplash.com/photo-1695048133142-1f6c7fdb9841?auto=format&fit=crop&w=900&q=80' },
  { id: 5, name: 'Sony WH-1000XM5', description: 'Премиум-наушники для путешествий', brand: 'Sony', category: 'Audio', image: 'https://images.unsplash.com/photo-1546435770-a3e426bf472b?auto=format&fit=crop&w=900&q=80' },
  { id: 6, name: 'Samsung Galaxy S24', description: 'Мощный смартфон с AI-функциями', brand: 'Samsung', category: 'Phones', image: 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?auto=format&fit=crop&w=900&q=80' },
  { id: 7, name: 'Apple Watch Ultra', description: 'Умные часы для спорта и приключений', brand: 'Apple', category: 'Wearables', image: 'https://images.unsplash.com/photo-1546868871-7041f2a55e12?auto=format&fit=crop&w=900&q=80' },
  { id: 8, name: 'Canon EOS R10', description: 'Компактная беззеркальная камера для фото и видео', brand: 'Canon', category: 'Photo', image: 'https://images.unsplash.com/photo-1516035069371-29a1b244cc32?auto=format&fit=crop&w=900&q=80' },
  { id: 9, name: 'Sony PlayStation 5', description: 'Игровая консоль нового поколения', brand: 'Sony', category: 'Gaming', image: 'https://images.unsplash.com/photo-1606813907291-d86efa9b94db?auto=format&fit=crop&w=900&q=80' },
  { id: 10, name: 'Dell XPS 15', description: 'Легкий ноутбук для работы и творчества', brand: 'Dell', category: 'Tech', image: 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?auto=format&fit=crop&w=900&q=80' },
  { id: 11, name: 'JBL Flip 6', description: 'Портативная колонка с мощным звуком', brand: 'JBL', category: 'Audio', image: 'https://images.unsplash.com/photo-1546435770-a3e426bf472b?auto=format&fit=crop&w=900&q=80' },
  { id: 12, name: 'GoPro Hero 12', description: 'Экшн-камера для путешествий и спорта', brand: 'GoPro', category: 'Photo', image: 'https://images.unsplash.com/photo-1526170375885-4d8ecf77b99f?auto=format&fit=crop&w=900&q=80' },
]

const imprints = ['#6d5efc', '#ff7a59', '#3cc7b7', '#f59e0b', '#ef4444', '#22c55e']

const productImageMap = {
  1: 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&w=900&q=80',
  2: 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?auto=format&fit=crop&w=900&q=80',
  3: 'https://images.unsplash.com/photo-1546435770-a3e426bf472b?auto=format&fit=crop&w=900&q=80',
  4: 'https://images.unsplash.com/photo-1695048133142-1f6c7fdb9841?auto=format&fit=crop&w=900&q=80',
  5: 'https://images.unsplash.com/photo-1546435770-a3e426bf472b?auto=format&fit=crop&w=900&q=80',
  6: 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?auto=format&fit=crop&w=900&q=80',
  7: 'https://images.unsplash.com/photo-1546868871-7041f2a55e12?auto=format&fit=crop&w=900&q=80',
  8: 'https://images.unsplash.com/photo-1516035069371-29a1b244cc32?auto=format&fit=crop&w=900&q=80',
  9: 'https://images.unsplash.com/photo-1606813907291-d86efa9b94db?auto=format&fit=crop&w=900&q=80',
  10: 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?auto=format&fit=crop&w=900&q=80',
  11: 'https://images.unsplash.com/photo-1546435770-a3e426bf472b?auto=format&fit=crop&w=900&q=80',
  12: 'https://images.unsplash.com/photo-1526170375885-4d8ecf77b99f?auto=format&fit=crop&w=900&q=80',
}

function getProductPrice(product) {
  const base = {
    1: 2499,
    2: 319,
    3: 289,
    4: 1299,
    5: 449,
    6: 1099,
  }

  return base[product.id] ?? 799 + (product.id % 5) * 57
}

function formatMoney(value) {
  return new Intl.NumberFormat('ru-RU', {
    style: 'currency',
    currency: 'RUB',
    maximumFractionDigits: 0,
  }).format(value)
}

function App() {
  const [products, setProducts] = useState(fallbackProducts)
  const [search, setSearch] = useState('')
  const [selectedCategory, setSelectedCategory] = useState('all')
  const [cart, setCart] = useState([])
  const [authMode, setAuthMode] = useState('login')
  const [authData, setAuthData] = useState({ phone: '', password: '' })
  const [user, setUser] = useState(null)
  const [statusMessage, setStatusMessage] = useState('')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const loadCatalog = async () => {
      try {
        const response = await fetch(`${API.catalog}/catalog`)
        if (!response.ok) {
          throw new Error('Catalog unavailable')
        }

        const data = await response.json()
        if (Array.isArray(data) && data.length > 0) {
          const enriched = [...data, ...fallbackProducts]
            .filter(
              (product, index, all) =>
                all.findIndex((item) => String(item.id) === String(product.id)) === index,
            )
            .map((product) => ({
              ...product,
              image:
                product.image ||
                productImageMap[product.id] ||
                productImageMap[product.name] ||
                fallbackProducts[(product.id || 1) % fallbackProducts.length].image,
              category: product.category || 'General',
              brand: product.brand || 'Marketplace',
            }))
            .slice(0, 12)

          setProducts(enriched)
        }
      } catch (error) {
        console.warn('Using fallback catalog data:', error)
      } finally {
        setLoading(false)
      }
    }

    loadCatalog()
  }, [])

  const categories = useMemo(
    () => ['all', ...new Set(products.map((product) => product.category).filter(Boolean))],
    [products],
  )

  const visibleProducts = products.filter((product) => {
    const matchesSearch = `${product.name} ${product.description} ${product.brand || ''}`
      .toLowerCase()
      .includes(search.toLowerCase())

    const matchesCategory =
      selectedCategory === 'all' || product.category === selectedCategory

    return matchesSearch && matchesCategory
  })

  const cartCount = cart.reduce((sum, item) => sum + item.quantity, 0)
  const totalAmount = cart.reduce((sum, item) => sum + item.price * item.quantity, 0)

  const addToCart = (product) => {
    setCart((current) => {
      const existing = current.find((item) => item.id === product.id)

      if (existing) {
        return current.map((item) =>
          item.id === product.id ? { ...item, quantity: item.quantity + 1 } : item,
        )
      }

      return [
        ...current,
        {
          id: product.id,
          name: product.name,
          description: product.description,
          category: product.category,
          price: getProductPrice(product),
          quantity: 1,
        },
      ]
    })
  }

  const updateQuantity = (productId, delta) => {
    setCart((current) =>
      current
        .map((item) =>
          item.id === productId ? { ...item, quantity: Math.max(0, item.quantity + delta) } : item,
        )
        .filter((item) => item.quantity > 0),
    )
  }

  const handleAuthSubmit = async (event) => {
    event.preventDefault()

    const endpoint = authMode === 'login' ? 'log' : 'reg'

    try {
      const response = await fetch(`${API.auth}/auth/${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          phone: authData.phone,
          password: authData.password,
        }),
      })

      const data = await response.json()

      if (!response.ok) {
        throw new Error(data.message || 'Ошибка авторизации')
      }

      if (authMode === 'login') {
        setUser({ phone: authData.phone, token: data.access_token })
        setStatusMessage('Вход выполнен успешно')
      } else {
        setStatusMessage('Регистрация успешна. Теперь можно войти в аккаунт.')
        setAuthMode('login')
      }

      setAuthData({ phone: '', password: '' })
    } catch (error) {
      setStatusMessage(error.message)
    }
  }

  const handleCheckout = () => {
    if (!user) {
      setStatusMessage('Чтобы оформить заказ, войдите в аккаунт')
      setAuthMode('login')
      return
    }

    if (cart.length === 0) {
      setStatusMessage('Корзина пуста')
      return
    }

    setStatusMessage(`Заказ на сумму ${formatMoney(totalAmount)} создан в демо-режиме`)
    setCart([])
  }

  return (
    <div className="marketplace-shell">
      <header className="topbar">
        <div className="brand-block">
          <div className="brand-mark">M</div>
          <div>
            <p className="eyebrow">Marketplace</p>
            <h1>Магазин для современного бизнеса</h1>
          </div>
        </div>

        <nav className="top-actions">
          <button type="button" className="menu-button muted">
            Каталог
          </button>
          <button type="button" className="menu-button muted">
            Доставка
          </button>
          <button type="button" className="menu-button accent">
            {user ? `Привет, ${user.phone}` : 'Войти'}
          </button>
        </nav>
      </header>

      <main className="marketplace-main">
        <section className="hero-banner">
          <div className="hero-copy">
            <span className="badge">Новая коллекция</span>
            <h2>Лучшие товары для работы, отдыха и жизни</h2>
            <p>
              Каталог собран из микросервисной архитектуры проекта: каталог товаров,
              корзина, пользователи и оформление заказов.
            </p>
            <div className="hero-stats">
              <div>
                <strong>12k+</strong>
                <span>товаров</span>
              </div>
              <div>
                <strong>4.9/5</strong>
                <span>оценка</span>
              </div>
              <div>
                <strong>24/7</strong>
                <span>поддержка</span>
              </div>
            </div>
          </div>

          <div className="hero-visual">
            <div className="floating-card card-a">Техно</div>
            <div className="floating-card card-b">Audio</div>
            <div className="floating-card card-c">Smart</div>
          </div>
        </section>

        <div className="content-grid">
          <section className="catalog-panel">
            <div className="catalog-header">
              <div>
                <p className="eyebrow">Каталог</p>
                <h3>Популярные товары</h3>
              </div>

              <label className="search-field">
                <span>Поиск</span>
                <input
                  type="text"
                  value={search}
                  onChange={(event) => setSearch(event.target.value)}
                  placeholder="Найти товар..."
                />
              </label>
            </div>

            <div className="category-bar">
              {categories.map((category) => (
                <button
                  key={category}
                  type="button"
                  className={selectedCategory === category ? 'category-pill active' : 'category-pill'}
                  onClick={() => setSelectedCategory(category)}
                >
                  {category === 'all' ? 'Все' : category}
                </button>
              ))}
            </div>

            {loading ? (
              <div className="state-box">Загрузка каталога...</div>
            ) : (
              <div className="product-grid">
                {visibleProducts.map((product) => {
                  const price = getProductPrice(product)
                  const cardColor = imprints[product.id % imprints.length]

                  return (
                    <article key={product.id} className="product-card">
                      <div className="product-thumb" style={{ background: `linear-gradient(135deg, ${cardColor}, #171c2a)` }}>
                        {product.image ? (
                          <img src={product.image} alt={product.name} />
                        ) : (
                          <span>{product.name.slice(0, 2).toUpperCase()}</span>
                        )}
                      </div>

                      <div className="product-body">
                        <div className="product-meta">
                          <span>{product.brand || 'Marketplace'}</span>
                          <span>{product.category || 'General'}</span>
                        </div>

                        <h4>{product.name}</h4>
                        <p>{product.description}</p>

                        <div className="product-action-row">
                          <strong>{formatMoney(price)}</strong>
                          <button type="button" onClick={() => addToCart(product)}>
                            В корзину
                          </button>
                        </div>
                      </div>
                    </article>
                  )
                })}
              </div>
            )}
          </section>

          <aside className="sidebar-panel">
            <div className="auth-card">
              <p className="eyebrow">Аккаунт</p>
              <h3>{user ? 'Профиль' : authMode === 'login' ? 'Вход' : 'Регистрация'}</h3>

              {user ? (
                <div className="user-box">
                  <strong>{user.phone}</strong>
                  <span>{user.token ? 'Token ready' : 'Guest mode'}</span>
                </div>
              ) : (
                <form onSubmit={handleAuthSubmit} className="auth-form">
                  <input
                    type="tel"
                    placeholder="Телефон"
                    value={authData.phone}
                    onChange={(event) =>
                      setAuthData((current) => ({ ...current, phone: event.target.value }))
                    }
                    required
                  />
                  <input
                    type="password"
                    placeholder="Пароль"
                    value={authData.password}
                    onChange={(event) =>
                      setAuthData((current) => ({ ...current, password: event.target.value }))
                    }
                    required
                  />

                  <button type="submit" className="submit-button">
                    {authMode === 'login' ? 'Войти' : 'Создать аккаунт'}
                  </button>
                </form>
              )}

              {!user && (
                <button
                  type="button"
                  className="switch-button"
                  onClick={() => setAuthMode(authMode === 'login' ? 'register' : 'login')}
                >
                  {authMode === 'login' ? 'Создать аккаунт' : 'Уже есть аккаунт'}
                </button>
              )}
            </div>

            <div className="cart-card">
              <div className="cart-header">
                <p className="eyebrow">Корзина</p>
                <span>{cartCount} товаров</span>
              </div>

              {cart.length === 0 ? (
                <div className="empty-state">Корзина пуста</div>
              ) : (
                <div className="cart-items">
                  {cart.map((item) => (
                    <div key={item.id} className="cart-item">
                      <div>
                        <h5>{item.name}</h5>
                        <p>{formatMoney(item.price)}</p>
                      </div>

                      <div className="stepper">
                        <button type="button" onClick={() => updateQuantity(item.id, -1)}>-</button>
                        <span>{item.quantity}</span>
                        <button type="button" onClick={() => updateQuantity(item.id, 1)}>+</button>
                      </div>
                    </div>
                  ))}
                </div>
              )}

              <div className="totals-row">
                <span>Итого</span>
                <strong>{formatMoney(totalAmount)}</strong>
              </div>

              <button type="button" className="checkout-button" onClick={handleCheckout}>
                Оформить заказ
              </button>
            </div>

            {statusMessage && <div className="notice">{statusMessage}</div>}
          </aside>
        </div>
      </main>
    </div>
  )
}

export default App
