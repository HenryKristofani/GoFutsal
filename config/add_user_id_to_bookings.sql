-- Add user_id column to bookings table for relationship with users
ALTER TABLE bookings ADD COLUMN user_id INTEGER REFERENCES users(id);

-- Optional: Create index for better query performance
CREATE INDEX idx_bookings_user_id ON bookings(user_id);