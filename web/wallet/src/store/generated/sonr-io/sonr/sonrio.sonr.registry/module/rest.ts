/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export interface RegistryMsgBuyAppAliasResponse {
  /** Did is the top level DID of the WhoIs. */
  did?: string;

  /** WhoIs is the updated WhoIs, contains the DID document and associated metadata. */
  who_is?: RegistryWhoIs;
}

export interface RegistryMsgBuyNameAliasResponse {
  /** Did is the top level DID of the WhoIs. */
  did?: string;

  /** WhoIs is the created WhoIs, contains the DID document and associated metadata. */
  who_is?: RegistryWhoIs;
}

export interface RegistryMsgCreateWhoIsResponse {
  /** Did is the top level DID of the created WhoIs. */
  did?: string;

  /** WhoIs is the created WhoIs, contains the DID document and associated metadata. */
  who_is?: RegistryWhoIs;
}

export interface RegistryMsgDeactivateWhoIsResponse {
  /** Success is true if the WhoIs was successfully deactivated. */
  success?: boolean;

  /** Did is the top level DID of the WhoIs. */
  did?: string;
}

export interface RegistryMsgTransferAppAliasResponse {
  /** Success is true if the Alias was successfully transferred. */
  success?: boolean;

  /** WhoIs is the updated WhoIs, contains the DID document and associated metadata. */
  who_is?: RegistryWhoIs;
}

export interface RegistryMsgTransferNameAliasResponse {
  /** Success is true if the Alias was successfully transferred. */
  success?: boolean;

  /** WhoIs is the updated WhoIs, contains the DID document and associated metadata. */
  who_is?: RegistryWhoIs;
}

export interface RegistryMsgUpdateWhoIsResponse {
  /** Did is the top level DID of the WhoIs. */
  did?: string;

  /** WhoIs is the created WhoIs, contains the DID document and associated metadata. */
  who_is?: RegistryWhoIs;
}

/**
 * Params defines the parameters for the module.
 */
export type RegistryParams = object;

export interface RegistryQueryAllWhoIsResponse {
  WhoIs?: RegistryWhoIs[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface RegistryQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: RegistryParams;
}

export interface RegistryQueryWhoIsAliasResponse {
  WhoIs?: RegistryWhoIs;
}

export interface RegistryQueryWhoIsControllerResponse {
  WhoIs?: RegistryWhoIs;
}

export interface RegistryQueryWhoIsResponse {
  WhoIs?: RegistryWhoIs;
}

export interface RegistryWhoIs {
  alias?: string[];

  /** Owner is the top level DID of the User or Application derived from the multisignature wallet. */
  owner?: string;

  /**
   * DIDDocument is the bytes representation of DIDDocument within the WhoIs. Initially marshalled as JSON.
   * @format byte
   */
  did_document?: string;
  controllers?: string[];
  type?: RegistryWhoIsType;

  /** @format int64 */
  timestamp?: string;
  is_active?: boolean;
}

export enum RegistryWhoIsType {
  USER = "USER",
  APPLICATION = "APPLICATION",
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title registry/v1/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryWhoIsAll
   * @summary Queries a list of WhoIs items.
   * @request GET:/sonr-io/sonr/registry/who_is
   */
  queryWhoIsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RegistryQueryAllWhoIsResponse, RpcStatus>({
      path: `/sonr-io/sonr/registry/who_is`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryWhoIs
   * @summary Queries a WhoIs by id.
   * @request GET:/sonr-io/sonr/registry/who_is/{did}
   */
  queryWhoIs = (did: string, params: RequestParams = {}) =>
    this.request<RegistryQueryWhoIsResponse, RpcStatus>({
      path: `/sonr-io/sonr/registry/who_is/${did}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryWhoIsAlias
   * @summary Queries a list of WhoIsAlias items.
   * @request GET:/sonr-io/sonr/registry/who_is_alias/{alias}
   */
  queryWhoIsAlias = (alias: string, params: RequestParams = {}) =>
    this.request<RegistryQueryWhoIsAliasResponse, RpcStatus>({
      path: `/sonr-io/sonr/registry/who_is_alias/${alias}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryWhoIsController
   * @summary Queries a list of WhoIsController items.
   * @request GET:/sonr-io/sonr/registry/who_is_controller/{controller}
   */
  queryWhoIsController = (controller: string, params: RequestParams = {}) =>
    this.request<RegistryQueryWhoIsControllerResponse, RpcStatus>({
      path: `/sonr-io/sonr/registry/who_is_controller/${controller}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/sonrio/sonr/registry/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<RegistryQueryParamsResponse, RpcStatus>({
      path: `/sonrio/sonr/registry/params`,
      method: "GET",
      format: "json",
      ...params,
    });
}
