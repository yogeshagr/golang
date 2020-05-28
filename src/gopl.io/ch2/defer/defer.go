package title

import "net/http"

func title(url string) error {
	resp, err := http.Get(url
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-type")
	if ct != "text/html" && !strings.HasPrefix(ct, "test/html;") {
		return fmt.Error("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// ...print dc's title element...

	return nil
}
