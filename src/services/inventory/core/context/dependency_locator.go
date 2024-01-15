package context

import "github.com/productivityeng/orabbit/db"


type DependencyLocator struct { 
	Client *db.PrismaClient
}

func NewDependencyLocator() (*DependencyLocator) { 
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	return &DependencyLocator{ Client: client }

}