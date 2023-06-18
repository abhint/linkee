class Linkee {
  constructor() {
    this.location = window.location.href;
    this.spinner = document.querySelector("#spinner");
    this.shortenBtn = document.querySelector("#shorten-btn");
    this.longUrlInputBox = document.querySelector("#long_url");
    this.shortUrlResultBox = document.querySelector("#short_url");
    this.shortUrlResultBlock = document.querySelector(
      "#short_url_result_block"
    );
    this.requestInProgress = false;
    this.previousUrl = "";

    this.shortenBtn.addEventListener("click", this.save.bind(this));
  }

  showLoading() {
    this.spinner.style.display = "inline-block";
    this.shortenBtn.innerHTML = `Loading...`;
    this.shortenBtn.disabled = true;
  }

  showUrl(r) {
    const shortenedUrl = `${this.location}${r.key}`;
    this.shortUrlResultBox.value = shortenedUrl;
    this.shortUrlResultBlock.style.display = "block";
    this.spinner.style.display = "none";
    this.shortenBtn.innerHTML = `Shorten`;
    this.shortenBtn.disabled = false;
    this.requestInProgress = false;
  }

  async request(url) {
    if (this.requestInProgress || url === this.previousUrl) {
      return;
    }

    this.showLoading();
    this.requestInProgress = true;
    this.previousUrl = url;

    try {
      const response = await fetch("/api", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ url }),
      });

      if (response.status !== 200) {
        this.shortenBtn.disabled = false;
      }

      const responseData = await response.json();
      this.showUrl(responseData);
    } catch (error) {
      console.error("Error:", error);
      this.requestInProgress = false;
    }
  }

  save() {
    let url = this.longUrlInputBox.value;
    if (!url) return;
    this.request(url);
  }
}

window.addEventListener("DOMContentLoaded", () => {
  const linkee = new Linkee();
  linkee.spinner.style.display = "none";
  linkee.shortenBtn.innerHTML = `Shorten`;
  linkee.shortUrlResultBlock.style.display = "none";
});
