// APIベースURL
const API_BASE_URL = 'http://localhost:8080';

// DOM要素の取得
const tabButtons = document.querySelectorAll('.tab-button');
const sections = document.querySelectorAll('.section');
const editModal = document.getElementById('edit-modal');
const editForm = document.getElementById('edit-form');
const closeModal = document.querySelector('.close');

// グローバル変数
let currentEditType = '';
let currentEditId = '';
let products = [];
let customers = [];
let orders = [];

// 初期化
document.addEventListener('DOMContentLoaded', function() {
    // タブ切り替えイベント
    tabButtons.forEach(button => {
        button.addEventListener('click', function() {
            switchTab(this.id.replace('-tab', ''));
        });
    });

    // フォーム送信イベント
    document.getElementById('product-form').addEventListener('submit', handleProductSubmit);
    document.getElementById('customer-form').addEventListener('submit', handleCustomerSubmit);
    document.getElementById('order-form').addEventListener('submit', handleOrderSubmit);
    document.getElementById('edit-form').addEventListener('submit', handleEditSubmit);

    // モーダル関連イベント
    closeModal.addEventListener('click', closeEditModal);
    document.getElementById('cancel-edit').addEventListener('click', closeEditModal);
    window.addEventListener('click', function(event) {
        if (event.target === editModal) {
            closeEditModal();
        }
    });

    // 初期データ読み込み
    loadAllData();
});

// タブ切り替え
function switchTab(tabName) {
    // タブボタンの状態更新
    tabButtons.forEach(button => {
        button.classList.remove('active');
    });
    document.getElementById(tabName + '-tab').classList.add('active');

    // セクションの表示切り替え
    sections.forEach(section => {
        section.classList.remove('active');
    });
    document.getElementById(tabName + '-section').classList.add('active');

    // データの読み込み
    switch(tabName) {
        case 'products':
            loadProducts();
            break;
        case 'customers':
            loadCustomers();
            break;
        case 'orders':
            loadOrders();
            loadProductsForSelect();
            loadCustomersForSelect();
            break;
    }
}

// 全データの読み込み
async function loadAllData() {
    await Promise.all([
        loadProducts(),
        loadCustomers(),
        loadOrders()
    ]);
}

// 商品一覧の読み込み
async function loadProducts() {
    try {
        const response = await fetch(`${API_BASE_URL}/products`);
        if (!response.ok) throw new Error('商品の読み込みに失敗しました');
        
        products = await response.json() || [];
        renderProductsTable();
    } catch (error) {
        console.error('Error loading products:', error);
        showMessage('error', error.message);
    }
}

// 顧客一覧の読み込み
async function loadCustomers() {
    try {
        const response = await fetch(`${API_BASE_URL}/customers`);
        if (!response.ok) throw new Error('顧客の読み込みに失敗しました');
        
        customers = await response.json() || [];
        await loadCustomerTotals();
        renderCustomersTable();
    } catch (error) {
        console.error('Error loading customers:', error);
        showMessage('error', error.message);
    }
}

// 顧客の合計金額を読み込み
async function loadCustomerTotals() {
    for (let customer of customers) {
        try {
            const response = await fetch(`${API_BASE_URL}/customers/${customer.id}/total`);
            if (response.ok) {
                const totalData = await response.json();
                customer.total = totalData.total || 0;
            } else {
                customer.total = 0;
            }
        } catch (error) {
            console.error(`Error loading total for customer ${customer.id}:`, error);
            customer.total = 0;
        }
    }
}

// 注文一覧の読み込み
async function loadOrders() {
    try {
        const response = await fetch(`${API_BASE_URL}/orders`);
        if (!response.ok) throw new Error('注文の読み込みに失敗しました');
        
        orders = await response.json() || [];
        renderOrdersTable();
    } catch (error) {
        console.error('Error loading orders:', error);
        showMessage('error', error.message);
    }
}

// 商品テーブルの描画
function renderProductsTable() {
    const tbody = document.querySelector('#products-table tbody');
    tbody.innerHTML = '';

    products.forEach(product => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${product.id}</td>
            <td>${product.name}</td>
            <td>¥${product.price.toLocaleString()}</td>
            <td>
                <button class="action-button edit" onclick="editProduct(${product.id})">編集</button>
                <button class="action-button delete" onclick="deleteProduct(${product.id})">削除</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 顧客テーブルの描画
function renderCustomersTable() {
    const tbody = document.querySelector('#customers-table tbody');
    tbody.innerHTML = '';

    customers.forEach(customer => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${customer.id}</td>
            <td>${customer.name}</td>
            <td>${customer.seat}</td>
            <td>¥${(customer.total || 0).toLocaleString()}</td>
            <td>
                <button class="action-button edit" onclick="editCustomer(${customer.id})">編集</button>
                <button class="action-button delete" onclick="deleteCustomer(${customer.id})">削除</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 注文テーブルの描画
function renderOrdersTable() {
    const tbody = document.querySelector('#orders-table tbody');
    tbody.innerHTML = '';

    orders.forEach(order => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${order.id}</td>
            <td>${order.product_id}</td>
            <td>${order.customer_id}</td>
            <td>${order.quantity}</td>
            <td>
                <button class="action-button edit" onclick="editOrder(${order.id})">編集</button>
                <button class="action-button delete" onclick="deleteOrder(${order.id})">削除</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 商品選択肢の読み込み
function loadProductsForSelect() {
    const select = document.getElementById('order-product');
    select.innerHTML = '<option value="">商品を選択</option>';
    
    products.forEach(product => {
        const option = document.createElement('option');
        option.value = product.id;
        option.textContent = `${product.name} (¥${product.price.toLocaleString()})`;
        select.appendChild(option);
    });
}

// 顧客選択肢の読み込み
function loadCustomersForSelect() {
    const select = document.getElementById('order-customer');
    select.innerHTML = '<option value="">顧客を選択</option>';
    
    customers.forEach(customer => {
        const option = document.createElement('option');
        option.value = customer.id;
        option.textContent = `${customer.name} (${customer.seat})`;
        select.appendChild(option);
    });
}

// 商品フォーム送信
async function handleProductSubmit(e) {
    e.preventDefault();
    
    const formData = new FormData(e.target);
    const product = {
        name: formData.get('name'),
        price: parseInt(formData.get('price'))
    };

    try {
        const response = await fetch(`${API_BASE_URL}/products`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(product)
        });

        if (!response.ok) throw new Error('商品の作成に失敗しました');

        showMessage('success', '商品が正常に追加されました');
        e.target.reset();
        loadProducts();
    } catch (error) {
        console.error('Error creating product:', error);
        showMessage('error', error.message);
    }
}

// 顧客フォーム送信
async function handleCustomerSubmit(e) {
    e.preventDefault();
    
    const formData = new FormData(e.target);
    const customer = {
        name: formData.get('name'),
        seat: formData.get('seat')
    };

    try {
        const response = await fetch(`${API_BASE_URL}/customers`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(customer)
        });

        if (!response.ok) throw new Error('顧客の作成に失敗しました');

        showMessage('success', '顧客が正常に追加されました');
        e.target.reset();
        loadCustomers();
    } catch (error) {
        console.error('Error creating customer:', error);
        showMessage('error', error.message);
    }
}

// 注文フォーム送信
async function handleOrderSubmit(e) {
    e.preventDefault();
    
    const formData = new FormData(e.target);
    const order = {
        product_id: parseInt(formData.get('product_id')),
        customer_id: parseInt(formData.get('customer_id')),
        quantity: parseInt(formData.get('quantity'))
    };

    try {
        const response = await fetch(`${API_BASE_URL}/orders`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(order)
        });

        if (!response.ok) throw new Error('注文の作成に失敗しました');

        showMessage('success', '注文が正常に追加されました');
        e.target.reset();
        loadOrders();
        loadCustomers(); // 顧客の合計金額を更新
    } catch (error) {
        console.error('Error creating order:', error);
        showMessage('error', error.message);
    }
}

// 編集フォーム送信
async function handleEditSubmit(e) {
    e.preventDefault();
    
    const formData = new FormData(e.target);
    let data = {};
    
    for (let [key, value] of formData.entries()) {
        if (key.includes('price') || key.includes('id') || key.includes('quantity')) {
            data[key] = parseInt(value);
        } else {
            data[key] = value;
        }
    }

    try {
        const response = await fetch(`${API_BASE_URL}/${currentEditType}/${currentEditId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) throw new Error('更新に失敗しました');

        showMessage('success', '正常に更新されました');
        closeEditModal();
        
        // データを再読み込み
        switch(currentEditType) {
            case 'products':
                loadProducts();
                break;
            case 'customers':
                loadCustomers();
                break;
            case 'orders':
                loadOrders();
                loadCustomers(); // 顧客の合計金額を更新
                break;
        }
    } catch (error) {
        console.error('Error updating:', error);
        showMessage('error', error.message);
    }
}

// 商品編集
function editProduct(id) {
    const product = products.find(p => p.id === id);
    if (!product) return;

    currentEditType = 'products';
    currentEditId = id;

    document.getElementById('modal-title').textContent = '商品編集';
    document.getElementById('edit-fields').innerHTML = `
        <div class="form-group">
            <label for="edit-name">商品名:</label>
            <input type="text" id="edit-name" name="name" value="${product.name}" required>
        </div>
        <div class="form-group">
            <label for="edit-price">価格:</label>
            <input type="number" id="edit-price" name="price" value="${product.price}" required>
        </div>
    `;

    editModal.style.display = 'block';
}

// 顧客編集
function editCustomer(id) {
    const customer = customers.find(c => c.id === id);
    if (!customer) return;

    currentEditType = 'customers';
    currentEditId = id;

    document.getElementById('modal-title').textContent = '顧客編集';
    document.getElementById('edit-fields').innerHTML = `
        <div class="form-group">
            <label for="edit-name">顧客名:</label>
            <input type="text" id="edit-name" name="name" value="${customer.name}" required>
        </div>
        <div class="form-group">
            <label for="edit-seat">座席:</label>
            <input type="text" id="edit-seat" name="seat" value="${customer.seat}" required>
        </div>
    `;

    editModal.style.display = 'block';
}

// 注文編集
function editOrder(id) {
    const order = orders.find(o => o.id === id);
    if (!order) return;

    currentEditType = 'orders';
    currentEditId = id;

    document.getElementById('modal-title').textContent = '注文編集';
    
    let productOptions = '<option value="">商品を選択</option>';
    products.forEach(product => {
        const selected = product.id === order.product_id ? 'selected' : '';
        productOptions += `<option value="${product.id}" ${selected}>${product.name} (¥${product.price.toLocaleString()})</option>`;
    });

    let customerOptions = '<option value="">顧客を選択</option>';
    customers.forEach(customer => {
        const selected = customer.id === order.customer_id ? 'selected' : '';
        customerOptions += `<option value="${customer.id}" ${selected}>${customer.name} (${customer.seat})</option>`;
    });

    document.getElementById('edit-fields').innerHTML = `
        <div class="form-group">
            <label for="edit-product">商品:</label>
            <select id="edit-product" name="product_id" required>
                ${productOptions}
            </select>
        </div>
        <div class="form-group">
            <label for="edit-customer">顧客:</label>
            <select id="edit-customer" name="customer_id" required>
                ${customerOptions}
            </select>
        </div>
        <div class="form-group">
            <label for="edit-quantity">数量:</label>
            <input type="number" id="edit-quantity" name="quantity" value="${order.quantity}" required min="1">
        </div>
    `;

    editModal.style.display = 'block';
}

// 削除機能
async function deleteProduct(id) {
    if (!confirm('この商品を削除しますか？')) return;

    try {
        const response = await fetch(`${API_BASE_URL}/products/${id}`, {
            method: 'DELETE'
        });

        if (!response.ok) throw new Error('商品の削除に失敗しました');

        showMessage('success', '商品が正常に削除されました');
        loadProducts();
    } catch (error) {
        console.error('Error deleting product:', error);
        showMessage('error', error.message);
    }
}

async function deleteCustomer(id) {
    if (!confirm('この顧客を削除しますか？')) return;

    try {
        const response = await fetch(`${API_BASE_URL}/customers/${id}`, {
            method: 'DELETE'
        });

        if (!response.ok) throw new Error('顧客の削除に失敗しました');

        showMessage('success', '顧客が正常に削除されました');
        loadCustomers();
    } catch (error) {
        console.error('Error deleting customer:', error);
        showMessage('error', error.message);
    }
}

async function deleteOrder(id) {
    if (!confirm('この注文を削除しますか？')) return;

    try {
        const response = await fetch(`${API_BASE_URL}/orders/${id}`, {
            method: 'DELETE'
        });

        if (!response.ok) throw new Error('注文の削除に失敗しました');

        showMessage('success', '注文が正常に削除されました');
        loadOrders();
        loadCustomers(); // 顧客の合計金額を更新
    } catch (error) {
        console.error('Error deleting order:', error);
        showMessage('error', error.message);
    }
}

// モーダルを閉じる
function closeEditModal() {
    editModal.style.display = 'none';
    currentEditType = '';
    currentEditId = '';
}

// メッセージ表示
function showMessage(type, message) {
    // 既存のメッセージを削除
    const existingMessage = document.querySelector('.message');
    if (existingMessage) {
        existingMessage.remove();
    }

    const messageDiv = document.createElement('div');
    messageDiv.className = `message ${type}`;
    messageDiv.textContent = message;

    const container = document.querySelector('.container');
    container.insertBefore(messageDiv, container.firstChild);

    // 3秒後に自動削除
    setTimeout(() => {
        messageDiv.remove();
    }, 3000);
}
