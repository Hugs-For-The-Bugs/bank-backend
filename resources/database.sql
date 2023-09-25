-- bank.sql

-- Create database
CREATE DATABASE IF NOT EXISTS bank;
USE bank;

-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
  id INT PRIMARY KEY AUTO_INCREMENT,
  first_name VARCHAR(50),
  surname VARCHAR(50),
  birthday VARCHAR(10),
  phone VARCHAR(15),
  email VARCHAR(50),
  ssn VARCHAR(14) UNIQUE,
  password VARCHAR(50),
  balance VARCHAR(50),
  active BOOLEAN
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

-- Insert data into account table
INSERT INTO accounts (first_name, surname, birthday, phone, email, ssn, password, balance, active) VALUES
('John', 'Doe', '19900101', '1234567890', 'john@email.com', '199001011234', 'password1', '1000', true),
('Jane', 'Doe', '19910202', '0987654321', 'jane@email.com', '199102022345', 'password2', '2000', true),
('Alice', 'Smith', '19850303', '1122334455', 'alice@email.com', '198503033456', 'password3', '3000', false),
('Bob', 'Johnson', '19800404', '2233445566', 'bob@email.com', '198004044567', 'password4', '4000', true),
('Tom', 'Cruise', '19620703', '9191919191', 'tom@email.com', '196207035678', 'password5', '5000', true);

-- Insert data into transaction table
INSERT INTO transactions (from_account_id, to_account_id, amount, type, state) VALUES
(1, 2, '100', 'Transfer', 'Successful'),
(2, 3, '50', 'Transfer', 'Failed'),
(3, 4, '200', 'Withdraw', 'Successful'),
(4, 5, '300', 'Deposit', 'Failed'),
(5, 1, '400', 'Transfer', 'Successful');

INSERT INTO admins (username, password) VALUES
('admin', 'admin');