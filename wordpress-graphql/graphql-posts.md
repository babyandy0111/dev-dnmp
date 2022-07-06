query postsQuery {
  posts {
    edges {
      node {
        title
        excerpt
        content
        slug
        seo {
          canonical
          metaDesc
          metaKeywords
          metaRobotsNofollow
          metaRobotsNoindex
          opengraphUrl
          opengraphTitle
          opengraphSiteName
          opengraphImage {
            altText
            fileSize
            mediaItemUrl
          }
        }
        author {
          node {
            email
            description
            firstName
            lastName
            slug
            uri
            seo {
              breadcrumbTitle
              metaDesc
              social {
                facebook
                instagram
                linkedIn
                mySpace
                pinterest
                soundCloud
                twitter
                wikipedia
                youTube
              }
            }
          }
        }
      }
    }
  }
}
