# stugo

**stugo** is a lightweight, file-based study web application written in Go. It allows users to revise subjects using Flashcards or Quizzes, loading content dynamically from JSON files.

## Features

- **No Database:** Data is loaded directly from local JSON files.
- **Two Modes:**
  - ⚡ **Flashcards:** Flip cards to memorize answers.
  - 📝 **Quiz:** Type the exact answer to test your knowledge.
- **Clean UI:** Built with [PicoCSS](https://picocss.com) for a minimal, responsive design.
- **Read-Only Web Interface:** Content is managed via the file system, not the browser.

## How to Run

1. **Prerequisites**: Ensure [Go](https://go.dev/dl/) is installed.
2. **Start the server**:
   ```bash
   go run main.go
   ```
3. **Open in Browser**:
   Visit [http://localhost:8080](http://localhost:8080).

## Adding Courses

To add a new subject, create a `.json` file inside the `data/` folder (e.g., `geography.json`).

**JSON Format:**

```json
{
  "title": "Capital Cities",
  "items": [
    {
      "question": "What is the capital of France?",
      "answer": "Paris"
    },
    {
      "question": "What is the capital of Japan?",
      "answer": "Tokyo"
    }
  ]
}
```

*Note: Changes to JSON files require a page refresh to appear in the app.*