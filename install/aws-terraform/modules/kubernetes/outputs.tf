output "cluster" {
  value = aws_eks_cluster.gitpod
}

output "cluster_auth" {
  value = data.aws_eks_cluster_auth.gitpod
}

output "vpc_id" {
  value = aws_vpc.gitpod.id
}

output "subnet_ids" {
  value = aws_subnet.public.*.id
}

output "worker_iam_role_arn" {
  value = aws_iam_role.node_group.arn
}