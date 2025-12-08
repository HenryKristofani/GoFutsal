-- Ensure username is unique in users table
-- This will add UNIQUE constraint if it doesn't exist
-- If it already exists, the constraint will be skipped

DO $$
BEGIN
    -- Try to add unique constraint
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint 
        WHERE conname = 'users_username_key'
    ) THEN
        ALTER TABLE users ADD CONSTRAINT users_username_key UNIQUE (username);
        RAISE NOTICE 'UNIQUE constraint added to username column';
    ELSE
        RAISE NOTICE 'UNIQUE constraint already exists on username column';
    END IF;
END $$;
