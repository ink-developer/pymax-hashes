import { ssrRenderAttrs } from "vue/server-renderer";
import { useSSRContext } from "vue";
import { _ as _export_sfc } from "./plugin-vue_export-helper.1tPrXgE0.js";
const __pageData = JSON.parse('{"title":"Формат данных","description":"","frontmatter":{},"headers":[],"relativePath":"data.md","filePath":"data.md"}');
const _sfc_main = { name: "data.md" };
function _sfc_ssrRender(_ctx, _push, _parent, _attrs, $props, $setup, $data, $options) {
  _push(`<div${ssrRenderAttrs(_attrs)}><h1 id="формат-данных" tabindex="-1">Формат данных <a class="header-anchor" href="#формат-данных" aria-label="Permalink to &quot;Формат данных&quot;">​</a></h1><p>Каждая версия содержит объект метаданных:</p><table tabindex="0"><thead><tr><th>Поле</th><th>Тип</th><th>Содержимое</th></tr></thead><tbody><tr><td><code>build_number</code></td><td>integer</td><td>Номер сборки</td></tr><tr><td><code>signature_scheme</code></td><td>string</td><td>Схема подписи APK</td></tr><tr><td><code>certificate_count</code></td><td>integer</td><td>Количество сертификатов</td></tr><tr><td><code>certificate_sha256</code></td><td>string[]</td><td>SHA-256-хеши сертификатов</td></tr><tr><td><code>certificate_meta_sha256</code></td><td>string</td><td>SHA-256 метаданных сертификатов</td></tr><tr><td><code>dex_meta_sha256</code></td><td>string</td><td>SHA-256 метаданных DEX-файлов</td></tr><tr><td><code>so_meta_sha256_arm64_v8a</code></td><td>string</td><td>SHA-256 метаданных нативных библиотек arm64-v8a</td></tr><tr><td><code>so_meta_sha256</code></td><td>object</td><td>SHA-256 метаданных нативных библиотек по архитектурам</td></tr></tbody></table><p>Объект <code>so_meta_sha256</code> содержит строковые поля <code>arm64-v8a</code>, <code>armeabi-v7a</code>, <code>x86</code> и <code>x86_64</code>.</p><p>Сервис хранит и возвращает переданные значения. Он не вычисляет хеши и не проверяет их соответствие APK. Алгоритм формирования полей <code>*_meta_sha256</code> в коде сервера не определён.</p><p>В <code>/versions.json</code> объекты метаданных сгруппированы по версии MAX. В <code>/versions/latest</code> тот же объект находится в поле <code>data</code>, а номер версии — в <code>version</code>.</p></div>`);
}
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("data.md");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
const data = /* @__PURE__ */ _export_sfc(_sfc_main, [["ssrRender", _sfc_ssrRender]]);
export {
  __pageData,
  data as default
};
