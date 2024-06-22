CREATE TABLE red_packets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    total_amount DECIMAL(10, 2) NOT NULL,
    total_count INT NOT NULL,
    remaining_amount DECIMAL(10, 2) NOT NULL,
    remaining_count INT NOT NULL
);

CREATE TABLE grabs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    red_packet_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    grabber VARCHAR(255) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (red_packet_id) REFERENCES red_packets(id)
);
