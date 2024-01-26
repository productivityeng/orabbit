"use server";

export type LockerType = "queue" | "user" | "exchange" | "virtualhost";

export type LockerModel = Awaited<ReturnType<typeof GetLocker>>;

export async function GetLocker(
  clusterId: number,
  lockerType: LockerType,
  artifactId: number
) {
  const importQueueFromClusterEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/locker/${lockerType}/${artifactId}`;

  const response = await fetch(importQueueFromClusterEndpoint, {
    method: "GET",
    cache: "no-store",
  });
  if (response.status != 200) {
    console.error(
      `Error fetching locker for ${lockerType} ${artifactId} from cluster ${clusterId}: ${
        response.statusText
      } with payload ${await response.text()}`
    );
    return null;
  }
  return (await response.json()) as {
    Id: number;
    Reason: string;
    Enabled: boolean;
    CreatedAt: Date;
    UpdatedAt: Date;
    UserResponsibleEmail: string;
  };
}

export async function CreateLockerAction(
  clusterId: number,
  lockerType: LockerType,
  artifactId: number,
  locker: { reason: string; responsible: string }
) {
  console.log(
    `Enviando locker ${JSON.stringify(locker)} para o cluster ${clusterId}`
  );
  const importQueueFromClusterEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/locker/${lockerType.toLowerCase()}/${artifactId}`;

  const response = await fetch(importQueueFromClusterEndpoint, {
    method: "POST",
    body: JSON.stringify(locker),
  });

  console.log(`Resposta do locker ${response.status}`);
  if (response.status != 201) {
    console.error(
      `Error creating locker for ${lockerType} ${artifactId} from cluster ${clusterId}: ${
        response.statusText
      } with payload ${await response.text()}`
    );
    return {
      ErrorMessage: `Error creating locker for ${lockerType} ${artifactId} from cluster ${clusterId}: ${await response.json()}`,
      Result: null,
    };
  }

  return { ErrorMessage: null, Result: (await response.json()) as LockManager };
}

export async function RemoveLockerAction(
  clusterId: number,
  lockerType: LockerType,
  lockerId: number
) {
  console.log(`Removendo locker do cluster ${clusterId}`);
  const removeLockerFromArtifactEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/locker/${lockerType.toLowerCase()}/${lockerId}`;

  console.log(`Sending request to ${removeLockerFromArtifactEndpoint}`);
  const response = await fetch(removeLockerFromArtifactEndpoint, {
    method: "DELETE",
    body: JSON.stringify({
      responsible: "vVictor",
    }),
    cache: "no-store",
  });
  if (response.status != 200) {
    console.error(
      `Error removing locker for ${lockerType} ${lockerId} from cluster ${clusterId}: ${
        response.statusText
      } with payload ${await response.text()}`
    );
    return {
      ErrorMessage: `Error creating locker for ${lockerType} with lockerId ${lockerId} from cluster ${clusterId}: ${await response.text()}`,
      Result: null,
    };
  }
  const result = (await response.json()) as LockManager;
  console.log(
    `Locker removido com sucesso do cluster ${clusterId} ${JSON.stringify(
      result
    )} | status code ${response.status}`
  );
  return {
    ErrorMessage: null,
    Result: result,
  };
}
