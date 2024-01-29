package db

import "context"




func (prisma *PrismaClient) GetClusterById(clusterId int,ctx context.Context) (cluster *ClusterModel,err error) {

	cluster,err = prisma.Cluster.FindUnique(Cluster.ID.Equals(clusterId)).Exec(ctx)
	return	
}