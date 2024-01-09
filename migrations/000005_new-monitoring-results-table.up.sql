CREATE TABLE IF NOT EXISTS monitoring_results (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), 
  url_id UUID NOT NULL REFERENCES urls(id) ON DELETE CASCADE,
  result VARCHAR(255) NOT NULL,
  status_code INTEGER NOT NULL,
  response_time INTEGER NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
