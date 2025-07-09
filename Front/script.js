//dont know if ill use this from ai but this is what i want to accoplish
document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("url-form");

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const longUrl = document.getElementById("long-url").value;

    try {
      const response = await fetch("http://localhost:8081/api/shorten", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ url: longUrl })
      });

      const data = await response.json();

      document.getElementById("result").innerHTML =
        // `<p>Short URL: <a href="${data.short_url}" target="_blank">${data.short_url}</a></p>`;
        `<p>Short URL: <a href="https://www.youtube.com/" target="_blank">${data.short_url}</a></p>`;
    } catch (err) {
      console.error("Error:", err);
      document.getElementById("result").textContent = "Failed to shorten URL.";
    }
  });
});

//leaving a note here to try reciving a long and small url disply the short and href be the long 
//but i dont think thats HOW they actually work i think its more when using a db and getting the short one based on the long one
// https://www.youtube.com/

