-- bank.sql

-- Create database
CREATE DATABASE IF NOT EXISTS bank;
USE bank;

-- Create accounts table with added datetime columns
CREATE TABLE IF NOT EXISTS accounts (
  id INT PRIMARY KEY AUTO_INCREMENT,
  first_name VARCHAR(50),
  surname VARCHAR(50),
  birthday VARCHAR(10),
  phone VARCHAR(15),
  email VARCHAR(50),
  social_security_number VARCHAR(14) UNIQUE,
  password VARCHAR(50),
  balance VARCHAR(50),
  active BOOLEAN,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
  id INT PRIMARY KEY AUTO_INCREMENT,
  from_account_id INT,
  to_account_id INT,
  amount VARCHAR(50),
  type ENUM('Transfer', 'Withdraw', 'Deposit'),
  state ENUM('Failed', 'Successful'),
  FOREIGN KEY (from_account_id) REFERENCES accounts(id),
  FOREIGN KEY (to_account_id) REFERENCES accounts(id)
);

-- Create admin table
CREATE TABLE IF NOT EXISTS admins (
  id INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(50),
  password VARCHAR(50)
);

-- Insert data into accounts table with datetime
INSERT INTO accounts (first_name, surname, birthday, phone, email, social_security_number, password, balance, active, created_at, updated_at) VALUES
('John', 'Doe', '19900101', '1234567890', 'john@email.com', '199001011234', '7c6a180b36896a0a8c02787eeafb0e4c', '1000', true, NOW(), NOW()),
('Jane', 'Doe', '19910202', '0987654321', 'jane@email.com', '199102022345', '6cb75f652a9b52798eb6cf2201057c73', '2000', true, NOW(), NOW()),
('Alice', 'Smith', '19850303', '1122334455', 'alice@email.com', '198503033456', '819b0643d6b89dc9b579fdfc9094f28e', '3000', false, NOW(), NOW()),
('Bob', 'Johnson', '19800404', '2233445566', 'bob@email.com', '198004044567', '34cc93ece0ba9e3f6f235d4af979b16c', '4000', true, NOW(), NOW()),
('Tom', 'Cruise', '19620703', '9191919191', 'tom@email.com', '196207035678', 'db0edd04aaac4506f7edab03ac855d56', '5000', true, NOW(), NOW());

-- Insert data into transactions table
INSERT INTO transactions (from_account_id, to_account_id, amount, type, state) VALUES
(1, 2, '100', 'Transfer', 'Successful'),
(2, 3, '50', 'Transfer', 'Failed'),
(3, 4, '200', 'Withdraw', 'Successful'),
(4, 5, '300', 'Deposit', 'Failed'),
(5, 1, '400', 'Transfer', 'Successful');

-- Insert data into admins table
INSERT INTO admins (username, password) VALUES
('admin', '21232f297a57a5a743894a0e4a801fc3');
