import { ssrRenderAttrs, ssrRenderStyle } from "vue/server-renderer";
import { useSSRContext } from "vue";
import { _ as _export_sfc } from "./plugin-vue_export-helper.1tPrXgE0.js";
const __pageData = JSON.parse('{"title":"PyMax Hashes","description":"","frontmatter":{},"headers":[],"relativePath":"index.md","filePath":"index.md"}');
const _sfc_main = { name: "index.md" };
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  _push(`<div${ssrRenderAttrs(_attrs)}><h1 id="pymax-hashes" tabindex="-1">PyMax Hashes <a class="header-anchor" href="#pymax-hashes" aria-label="Permalink to &quot;PyMax Hashes&quot;">​</a></h1><p>Сервис метаданных для разных версий MAX, дополняющий <a href="https://github.com/MaxApiTeam/PyMax" target="_blank" rel="noreferrer">PyMax</a>. Хранит номер сборки, схему подписи и SHA-256-хеши сертификатов, DEX-файлов и нативных библиотек.</p><p>Чтение данных доступно без авторизации:</p><div class="language-sh vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">sh</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="${ssrRenderStyle({ "--shiki-light": "#6F42C1", "--shiki-dark": "#B392F0" })}">curl</span><span style="${ssrRenderStyle({ "--shiki-light": "#032F62", "--shiki-dark": "#9ECBFF" })}"> https://hashes.pymax.org/versions.json</span></span></code></pre></div><ul><li><a href="./api.html">API</a> — запросы, ответы и ошибки.</li><li><a href="./data.html">Формат данных</a> — поля метаданных.</li><li><a href="./self-hosting.html">Самостоятельный запуск</a> — Docker Compose и настройки.</li></ul></div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("index.md");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const index = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender]]);
export {
  __pageData,
  index as default
};
