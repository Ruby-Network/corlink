export default defineAppConfig({
  docus: {
    title: 'Corlink',
    description: 'The best place to start your documentation.',
    image: 'https://vercel.com/favicon.ico',
    socials: {
      github: 'ruby-network/corlink',
    },
    github: {
      dir: 'docs/content',
      branch: 'main',
      repo: 'corlink',
      owner: 'ruby-network',
      edit: true
    },
    aside: {
      level: 0,
      collapsed: false,
      exclude: []
    },
    main: {
      padded: true,
      fluid: true
    },
    header: {
      title: 'Corlink',
      showLinkIcon: true,
      exclude: [],
      fluid: true
    },
    footer: {
        credits: {
            icon: '',
            text: 'Created By Ruby Network',
            href: 'https://rubynetwork.co'
        },
    }
  }
})
