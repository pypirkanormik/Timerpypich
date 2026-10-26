package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, htmlPage)
	})
	http.HandleFunc("/scream.mp3", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "scream.mp3")
	})
	http.HandleFunc("/scare.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "scare.png")
	})

	addr := ":3000"
	fmt.Printf("\n  PYPЫRkA life\n")
	fmt.Printf("  http://localhost%s\n\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

const htmlPage = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
<title>PYPЫRkA life</title>
<style>
  @import url('https://fonts.googleapis.com/css2?family=Share+Tech+Mono&display=swap');

  * { margin: 0; padding: 0; box-sizing: border-box; }

  body {
    background: #0a0a0a;
    color: #e8e8e8;
    font-family: 'Share Tech Mono', monospace;
    min-height: 100vh;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    position: relative;
  }

  body::before {
    content: '';
    position: fixed;
    inset: 0;
    background: radial-gradient(ellipse at center, transparent 25%, rgba(0,0,0,0.75) 100%);
    pointer-events: none;
    z-index: 2;
  }

  .container {
    position: relative;
    z-index: 10;
    text-align: center;
    padding: 20px;
  }

  .cat-wrap {
    perspective: 800px;
    margin-bottom: 1.4rem;
    height: 90px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
  }

  .cat {
    font-size: 4.2rem;
    line-height: 1;
    transform-style: preserve-3d;
    animation: catSpin 14s linear infinite;
    filter: grayscale(1) brightness(1.35) drop-shadow(0 0 12px rgba(255,255,255,0.25));
    user-select: none;
  }

  @keyframes catSpin {
    0%   { transform: rotateY(0deg); }
    100% { transform: rotateY(360deg); }
  }

  h1 {
    font-size: clamp(2.4rem, 10vw, 4.5rem);
    letter-spacing: 0.12em;
    color: #f5f5f5;
    text-shadow: 0 0 25px rgba(255,255,255,0.2);
    margin-bottom: 2.8rem;
    font-weight: 400;
  }

  .timer {
    display: flex;
    justify-content: center;
    gap: 0.5rem;
    flex-wrap: wrap;
  }

  .unit {
    display: flex;
    flex-direction: column;
    align-items: center;
    min-width: 68px;
  }

  .digits {
    font-size: clamp(2rem, 8vw, 3.2rem);
    color: #ffffff;
    letter-spacing: 0.05em;
    line-height: 1;
    text-shadow: 0 0 14px rgba(255,255,255,0.3);
  }

  .label {
    font-size: 0.65rem;
    color: #777;
    letter-spacing: 0.25em;
    margin-top: 0.5rem;
    text-transform: uppercase;
  }

  .sep {
    font-size: 2.2rem;
    color: #555;
    align-self: center;
    padding-bottom: 1rem;
    animation: blink 1.2s step-end infinite;
  }

  @keyframes blink {
    0%, 100% { opacity: 1; }
    50% { opacity: 0.2; }
  }

  .word {
    position: fixed;
    top: -50px;
    color: rgba(255,255,255,0.55);
    font-size: 0.9rem;
    letter-spacing: 0.18em;
    text-transform: uppercase;
    white-space: nowrap;
    pointer-events: none;
    z-index: 1;
    text-shadow: 0 0 8px rgba(255,255,255,0.15);
    will-change: transform, opacity;
  }

  .footer {
    position: fixed;
    bottom: 14px;
    left: 0;
    width: 100%;
    text-align: center;
    font-size: 0.6rem;
    color: #444;
    letter-spacing: 0.2em;
    z-index: 20;
    cursor: pointer;
    transition: color 0.2s;
  }

  .footer:hover { color: #888; }

  .scan {
    position: fixed;
    inset: 0;
    background: repeating-linear-gradient(
      0deg,
      transparent,
      transparent 2px,
      rgba(0,0,0,0.04) 2px,
      rgba(0,0,0,0.04) 4px
    );
    pointer-events: none;
    z-index: 3;
  }

  /* SCREAMER */
  #screamer {
    display: none;
    position: fixed;
    inset: 0;
    background: #000;
    z-index: 9999;
    overflow: hidden;
  }

  #screamer.on {
    display: block;
  }

  #screamer img {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 100%;
    height: 100%;
    object-fit: cover;
    transform: translate(-50%, -50%) scale(1.1);
    animation: scareShake 0.07s infinite, scareFlash 0.12s infinite;
    filter: contrast(1.4) brightness(1.1) saturate(1.3);
  }

  #screamer .veil {
    position: absolute;
    inset: 0;
    background: radial-gradient(circle, transparent 30%, rgba(0,0,0,0.55) 100%);
    pointer-events: none;
    z-index: 2;
  }

  #screamer .lines {
    position: absolute;
    inset: 0;
    background: repeating-linear-gradient(
      0deg,
      transparent,
      transparent 3px,
      rgba(0,0,0,0.25) 3px,
      rgba(0,0,0,0.25) 5px
    );
    pointer-events: none;
    z-index: 3;
    animation: linesMove 0.1s steps(2) infinite;
  }

  @keyframes scareShake {
    0%   { transform: translate(-50%, -50%) scale(1.12) rotate(0deg); }
    25%  { transform: translate(-49%, -51%) scale(1.18) rotate(-1.5deg); }
    50%  { transform: translate(-51%, -49%) scale(1.08) rotate(1.2deg); }
    75%  { transform: translate(-48%, -50%) scale(1.2) rotate(-0.8deg); }
    100% { transform: translate(-50%, -50%) scale(1.12) rotate(0deg); }
  }

  @keyframes scareFlash {
    0%, 100% { filter: contrast(1.4) brightness(1.1); }
    40% { filter: contrast(2) brightness(0.35) hue-rotate(20deg); }
    70% { filter: contrast(1.8) brightness(1.6); }
  }

  @keyframes linesMove {
    0% { transform: translateY(0); }
    100% { transform: translateY(4px); }
  }
</style>
</head>
<body>
  <div class="scan"></div>
  <div class="container">
    <div class="cat-wrap" id="catBtn">
      <div class="cat">🐱</div>
    </div>
    <h1>PYPЫRkA life</h1>
    <div class="timer" id="timer">
      <div class="unit">
        <div class="digits" id="days">00</div>
        <div class="label">days</div>
      </div>
      <div class="sep">:</div>
      <div class="unit">
        <div class="digits" id="hours">00</div>
        <div class="label">hours</div>
      </div>
      <div class="sep">:</div>
      <div class="unit">
        <div class="digits" id="mins">00</div>
        <div class="label">mins</div>
      </div>
      <div class="sep">:</div>
      <div class="unit">
        <div class="digits" id="secs">00</div>
        <div class="label">secs</div>
      </div>
    </div>
  </div>
  <div class="footer" id="footerLink">pypirka™ unionteams.ru</div>

  <div id="screamer">
    <img src="/scare.png" alt="" id="scareImg" decoding="async">
    <div class="veil"></div>
    <div class="lines"></div>
  </div>

  <audio id="screamAudio" preload="auto" src="/scream.mp3"></audio>

<script>
  const scareImg = document.getElementById('scareImg');
  const screamAudio = document.getElementById('screamAudio');
  if (scareImg.decode) scareImg.decode().catch(()=>{});
  screamAudio.load();

  const target = new Date('2026-12-01T00:00:00').getTime();
  const words = [
    'suicide', 'death', 'betrayal', 'betrayal',
    'suicide', 'death', 'emptiness', 'regret',
    'void', 'alone', 'end', 'nothing',
    'decay', 'fade', 'silence', 'loss',
    'hollow', 'pain', 'gone', 'broken'
  ];

  function pad(n) {
    return n.toString().padStart(2, '0');
  }

  function update() {
    const now = Date.now();
    let diff = target - now;
    if (diff <= 0) {
      ['days','hours','mins','secs'].forEach(id => document.getElementById(id).textContent = '00');
      return;
    }
    const days = Math.floor(diff / 86400000); diff %= 86400000;
    const hours = Math.floor(diff / 3600000); diff %= 3600000;
    const mins = Math.floor(diff / 60000); diff %= 60000;
    const secs = Math.floor(diff / 1000);
    document.getElementById('days').textContent = pad(days);
    document.getElementById('hours').textContent = pad(hours);
    document.getElementById('mins').textContent = pad(mins);
    document.getElementById('secs').textContent = pad(secs);
  }
  update();
  setInterval(update, 1000);

  function spawnWord() {
    const el = document.createElement('div');
    el.className = 'word';
    el.textContent = words[Math.floor(Math.random() * words.length)];
    const startX = Math.random() * 100;
    const drift = (Math.random() - 0.5) * 40;
    const duration = 5 + Math.random() * 7;
    const size = 0.75 + Math.random() * 0.9;
    const startOpacity = 0.25 + Math.random() * 0.45;
    const rotate = (Math.random() - 0.5) * 30;
    el.style.left = startX + 'vw';
    el.style.fontSize = size + 'rem';
    el.style.opacity = startOpacity;
    const start = performance.now();
    function animate(now) {
      const t = (now - start) / (duration * 1000);
      if (t >= 1) { el.remove(); return; }
      const y = t * t * 115;
      const x = startX + Math.sin(t * Math.PI * 2) * (drift * 0.3) + drift * t;
      el.style.transform = 'translateY(' + y + 'vh) translateX(' + (x - startX) + 'vw) rotate(' + (rotate * t) + 'deg)';
      el.style.opacity = startOpacity * (1 - t * 0.85);
      requestAnimationFrame(animate);
    }
    requestAnimationFrame(animate);
    document.body.appendChild(el);
  }
  setInterval(spawnWord, 280);

  document.getElementById('footerLink').addEventListener('click', () => {
    window.location.href = 'https://unionteams.ru/player';
  });

  document.getElementById('catBtn').addEventListener('click', triggerScreamer);

  function triggerScreamer() {
    const sc = document.getElementById('screamer');
    sc.classList.add('on');

    screamAudio.currentTime = 0;
    screamAudio.volume = 1;
    const p = screamAudio.play();
    if (p && p.catch) p.catch(() => {});

    const end = () => {
      document.body.innerHTML = '';
      document.body.style.background = '#000';
      try { window.close(); } catch(e) {}
      setTimeout(() => { window.location.href = 'about:blank'; }, 150);
    };

    screamAudio.onended = end;
    setTimeout(end, 3000);
  }
</script>
</body>
</html>`
