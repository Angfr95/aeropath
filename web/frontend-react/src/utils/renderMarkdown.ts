export function renderMarkdown(text: string): string {
  if (!text) return "";
  let html = text
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;");
  // Keep links and local illustrations out of the formatting replacements.
  // Raw HTML remains escaped; images may only use our illustration directory.
  const embeds: string[] = [];
  const attribute = (value: string) => value.replace(/"/g, "&quot;").replace(/'/g, "&#39;");
  const embed = (value: string) => {
    const token = `%AEROPATH_EMBED_${embeds.length}%`;
    embeds.push(value);
    return token;
  };
  html = html.replace(/!\[([^\]\n]*)\]\((\/illustrations\/[a-zA-Z0-9/_-]+\.svg)\)/g,
    (_, alt: string, src: string) => embed(`<img src="${src}" alt="${attribute(alt)}" loading="lazy" class="w-full h-auto rounded-xl my-6" />`));
  html = html.replace(/\[([^\]\n]+)\]\((https:\/\/[^\s"'<>\\()]+)\)/g,
    (_, label: string, href: string) => embed(`<a href="${attribute(href)}" class="text-blue-400 underline hover:text-blue-300">${label}</a>`));
  html = html.replace(/^### (.+)$/gm, '<h3 class="text-lg font-bold text-white mt-4 mb-2">$1</h3>');
  html = html.replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold text-white mt-5 mb-2">$1</h2>');
  html = html.replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold text-white mt-6 mb-3">$1</h1>');
  html = html.replace(/\*\*(.+?)\*\*/g, '<strong class="text-white font-bold">$1</strong>');
  html = html.replace(/\*(.+?)\*/g, '<em class="text-slate-200 italic">$1</em>');
  html = html.replace(/^- (.+)$/gm, '<li class="text-slate-300 ml-4 list-disc">$1</li>');
  html = html.replace(/^\d+\. (.+)$/gm, '<li class="text-slate-300 ml-4 list-decimal">$1</li>');
  html = html.replace(/\n\n/g, '</p><p class="text-slate-300 mb-3 leading-relaxed">');
  html = '<p class="text-slate-300 mb-3 leading-relaxed">' + html + '</p>';
  return html.replace(/%AEROPATH_EMBED_(\d+)%/g, (token, index: string) => embeds[Number(index)] ?? token);
}
