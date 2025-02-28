particlesJS("particles-js", {
  particles: {
    number: { value: 80, density: { enable: true, value_area: 800 } },
    color: { value: "#1a73e8" },
    shape: { type: "circle" },
    opacity: {
      value: 0.5,
      random: false,
      animation: { enable: true, speed: 1, minimumValue: 0.1, sync: false }
    },
    size: {
      value: 3,
      random: true,
      animation: { enable: true, speed: 2, minimumValue: 0.1, sync: false }
    },
    lineLinked: {
      enable: true,
      distance: 150,
      color: "#1a73e8",
      opacity: 0.4,
      width: 1
    },
    move: {
      enable: true,
      speed: 2,
      direction: "none",
      random: false,
      straight: false,
      outMode: "bounce",
      attract: { enable: false, rotateX: 600, rotateY: 1200 }
    }
  },
  interactivity: {
    detectOn: "canvas",
    events: {
      onHover: { enable: true, mode: "grab" },
      onClick: { enable: true, mode: "push" },
      resize: true
    },
    modes: {
      grab: { distance: 140, lineLinked: { opacity: 1 } },
      push: { particles_nb: 4 }
    }
  },
  retina_detect: true
});