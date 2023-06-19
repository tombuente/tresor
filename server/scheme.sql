CREATE TABLE IF NOT EXISTS code_languages (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS code_snippets (
	id INTEGER PRIMARY KEY,
	content TEXT NOT NULL,
	language_id INTEGER NOT NULL,
	FOREIGN KEY (language_id) REFERENCES code_languages (id)
);
