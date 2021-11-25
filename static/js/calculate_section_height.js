
var nav_h = document.querySelector("nav").offsetHeight;
var foo_h = document.querySelector("footer").offsetHeight;

var viewport_h = Math.max(document.documentElement.clientHeight || 0, window.innerHeight || 0);

var content_section = document.getElementById("content");
var content_section_margin_top = parseInt(window.getComputedStyle(content_section).marginTop);

var content_section_h = viewport_h - (nav_h + foo_h + content_section_margin_top);

var current_content_section_h = parseInt(window.getComputedStyle(content_section).height);

if (current_content_section_h < content_section_h) {
    content_section.style.height = content_section_h + "px";
}
