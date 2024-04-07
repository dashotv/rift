// Code generated by github.com/dashotv/golem. DO NOT EDIT.

export interface Page {
  id?: string;
  created_at?: string;
  updated_at?: string;

  name?: string;
  url?: string;
  scraper?: string;
  downloader?: string;
  processed_at?: string;
}

export interface Video {
  id?: string;
  created_at?: string;
  updated_at?: string;

  page_id?: string;
  title?: string;
  season?: number;
  episode?: number;
  raw?: string;
  display_id?: string;
  extension?: string;
  resolution?: number;
  size?: number;
  download?: string;
  view?: string;
  source?: string;
}

export interface Visit {
  id?: string;
  created_at?: string;
  updated_at?: string;

  page_id?: string;
  url?: string;
  error?: string;
  stacktrace?: string[];
}
