# profolio-link-replacer

A tiny CLI tool to update links in your Markdown files to use Hugo shortcodes compatible with the [`hugoprofolio`](https://github.com/0hlov3/hugo-theme-hugoprofolio) theme.

If you wrote your draft Markdown in any tool (like Notion, Obsidian, etc.), this helps you convert regular Markdown links into the theme's `{{< newtablink >}}` shortcode format.

## 🔧 What It Does

It replaces links like:

```markdown
[Prometheus](<> "Prometheus")
[Loki](<https://grafana.com/oss/loki/> "Loki")
```

with

```markdown
{{< newtablink "https://prometheus.io/" >}}Prometheus{{< /newtablink >}}
{{< newtablink "https://grafana.com/oss/loki/" >}}Loki{{< /newtablink >}}
```

## Usage
```shell
go run main.go your_markdown_file.md
```

## Why?
The hugoprofolio theme uses shortcodes for external links that open in a new tab. This tool makes it easy to prepare your Markdown files without manually editing every link.

## TODO
- Optional dry-run mode
- Directory support
- More robust fallback link mapping
