import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'CLI',
  description: 'A simple and powerful command line framework for Go',
  base: '/cli/',
  
  head: [
    ['link', { rel: 'icon', href: '/cli/favicon.ico' }]
  ],

  themeConfig: {
    logo: '/logo.svg',
    
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Getting Started', link: '/guide/getting-started' },
      { text: 'Guide', link: '/guide/' },
      { text: 'Examples', link: '/examples/' },
      { text: 'API', link: '/api/' },
      {
        text: 'Links',
        items: [
          { text: 'GitHub', link: 'https://github.com/go-zoox/cli' },
          { text: 'PkgGoDev', link: 'https://pkg.go.dev/github.com/go-zoox/cli' }
        ]
      }
    ],

    sidebar: {
      '/guide/': [
        {
          text: 'Introduction',
          items: [
            { text: 'Getting Started', link: '/guide/getting-started' },
            { text: 'Installation', link: '/guide/installation' }
          ]
        },
        {
          text: 'Basic Usage',
          items: [
            { text: 'Single Command', link: '/guide/single-command' },
            { text: 'Multiple Commands', link: '/guide/multiple-commands' },
            { text: 'Flags', link: '/guide/flags' }
          ]
        },
        {
          text: 'Interactive Components',
          items: [
            { text: 'Text Input', link: '/guide/interactive/text' },
            { text: 'Select', link: '/guide/interactive/select' },
            { text: 'Confirm', link: '/guide/interactive/confirm' },
            { text: 'Password', link: '/guide/interactive/password' },
            { text: 'Multiselect', link: '/guide/interactive/multiselect' }
          ]
        },
        {
          text: 'Loading Components',
          items: [
            { text: 'Spinner', link: '/guide/loading/spinner' },
            { text: 'Progress', link: '/guide/loading/progress' }
          ]
        },
        {
          text: 'Advanced',
          items: [
            { text: 'Daemon Mode', link: '/guide/daemon' }
          ]
        }
      ],
      '/examples/': [
        {
          text: 'Examples',
          items: [
            { text: 'Single Command', link: '/examples/single' },
            { text: 'Multiple Commands', link: '/examples/multiple' },
            { text: 'Interactive', link: '/examples/interactive' },
            { text: 'Loading', link: '/examples/loading' }
          ]
        }
      ],
      '/api/': [
        {
          text: 'API Reference',
          items: [
            { text: 'Overview', link: '/api/' },
            { text: 'SingleProgram', link: '/api/single-program' },
            { text: 'MultipleProgram', link: '/api/multiple-program' },
            { text: 'Interactive', link: '/api/interactive' },
            { text: 'Loading', link: '/api/loading' }
          ]
        }
      ]
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/go-zoox/cli' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2024-present go-zoox'
    },

    search: {
      provider: 'local'
    }
  }
})
