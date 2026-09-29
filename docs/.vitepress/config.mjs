import { defineConfig } from 'vitepress'

export default defineConfig({
  lang: 'ru-RU',
  title: 'PyMax Hashes',
  description: 'Метаданные и SHA-256-хеши для версий MAX.',
  themeConfig: {
    nav: [
      { text: 'Документация', link: '/' },
      { text: 'API', link: '/api' }
    ],
    sidebar: [
      { text: 'Введение', link: '/' },
      { text: 'API', link: '/api' },
      { text: 'Формат данных', link: '/data' },
      { text: 'Самостоятельный запуск', link: '/self-hosting' }
    ],
    outline: { label: 'На этой странице' },
    docFooter: { prev: 'Назад', next: 'Далее' },
    sidebarMenuLabel: 'Меню',
    returnToTopLabel: 'Наверх'
  }
})
