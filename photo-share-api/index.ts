const {ApolloServer} = require('apollo-server')
const {GraphQLScalarType} = require('graphql')

var users = [
    { "githubLogin": "mHattrup", "name": "Mike Hattrup" },
    { "githubLogin": "gPlake", "name": "Glen Plake" },
    { "githubLogin": "sSchmidt", "name": "Scot Schmidt" }
]

var tags = [
    { "photoID": "1", "userID": "gPlake" },
    { "photoID": "2", "userID": "sSchmidt" },
    { "photoID": "2", "userID": "mHattrup" },
    { "photoID": "2", "userID": "gPlake" }
]

var photos = [
    {
        "id": "1",
        "name": "Dropping the Heart Chute",
        "description": "The heart chute is one of my favorite chutes",
        "category": "ACTION",
        "githubUser": "gPlake",
        "created": "1-2-1985"
    },
    {
        "id": "2",
        "name": "Enjoying the sunshine",
        "category": "SELFIE",
        "githubUser": "sSchmidt",
        "created": "3-28-1977"
    },
    {
        "id": "3",
        "name": "Gunbarrel 25",
        "description": "25 laps on gunbarrel today",
        "category": "LANDSCAPE",
        "githubUser": "sSchmidt",
        "created": "2018-04-15T19:09:57.308Z"
    }
]
let id: number = 0;



const typeDefs = `

    scalar DateTime
    
    enum PhotoCategory {
    SELFIE
    PORTRAIT
    ACTION
    LANDSCAPE
    GRAPHIC
    }
    
    type Photo {
      id: ID!
      url: String!
      name: String!
      description: String
      category: PhotoCategory
      postedBy: User!
      taggedUsers:[User!]!
      created: DateTime!
    }
    
    type User {
        githubLogin: ID!
        name: String
        avatar: String
        postedPhotos: [Photo!]!
        inPhotos:[Photo!]!
    }
    
    
    
    
    input PostPhotoInput {
    name: String!,
    description: String,
    category: PhotoCategory=GRAPHIC
    }
    
    type Query {
        totalPhotos: Int!
        allPhotos: [Photo!]!
        allUsers:[User]!
    }
    type Mutation {
    postPhoto(input:PostPhotoInput): Photo!
    }
 `
const resolvers = {
    Query: {
        totalPhotos: () => photos.length,
        allPhotos: () => photos,
        allUsers: () =>users
    },
    Mutation: {
        postPhoto(parent, args) {
            let newPhoto = {
                id: id++,
                ...args.input,
                created: new Date()
            }
            photos.push(newPhoto)
            photos.forEach(pic => console.log(pic))
            return newPhoto;
        }

    },
    Photo: {
        url: parent => `http://yoursite.com/img/${parent.id}.jpg`,
        postedBy: parent => { return users.find(user => user.githubLogin === parent.githubUser) },
        taggedUsers: parent =>  tags.filter(tag => tag.photoID === parent.id).map(tag => tag.userID).map(userid => users.find(u => u.githubLogin === userid))
    },
    User: {
        postedPhotos: parent => {
            return photos.filter(p => p.githubUser === parent.githubLogin)
        },
        inPhotos: parent => {
            return tags.filter(tag => tag.userID === parent.userID).map(tag => tag.photoID).map(photoId => photos.find(photo => photo.id === photoId))
        }
    },
    DateTime: new GraphQLScalarType({
        name: 'DateTime',
        description: ' A valid date',
        parseValue: value => new Date(value),
        serialize: value => new Date(value).toISOString(),
        parseLiteral: ast => ast.value
    })
}

// 2. Create a new instance of the server.
// 3. Send it an object with typeDefs (the schema) and resolvers
const server = new ApolloServer({
    typeDefs,
    resolvers
})

server.listen().then(({url}) => console.log(`GraphQL Service running on ${url}`))