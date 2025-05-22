-- Table: students
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(190) DEFAULT '',
    last_name VARCHAR(190) DEFAULT '',
    email VARCHAR(120) UNIQUE DEFAULT '',
    phone VARCHAR(120) UNIQUE DEFAULT '',
    password VARCHAR(255) DEFAULT '',
    avatar VARCHAR(190) DEFAULT '',
    address VARCHAR(500) DEFAULT '',
    validated BOOLEAN DEFAULT TRUE,
    fcmtoken VARCHAR(190) DEFAULT '',
    token VARCHAR(190) UNIQUE DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Table: wallet
CREATE TABLE wallet (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    balance DECIMAL(10,2) DEFAULT 0.00,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id),
    INDEX idx_wallet_student_id (student_id)
);

-- Table: recharge_cards
CREATE TABLE recharge_cards (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) UNIQUE NOT NULL,
    value DECIMAL(11,0) NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    used_at TIMESTAMP NULL,
    used_by_id INT NULL,
    FOREIGN KEY (used_by_id) REFERENCES students(id)
);

-- Table: orders add some column 
ALTER TABLE orders
ADD COLUMN payment_method VARCHAR(255),
ADD COLUMN student_id INT NOT NULL,
ADD COLUMN order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE orders
ADD CONSTRAINT fk_orders_student
FOREIGN KEY (student_id) REFERENCES students(id);
